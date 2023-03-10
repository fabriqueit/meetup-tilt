# version_settings() enforces a minimum Tilt version
# https://docs.tilt.dev/api.html#api.version_settings
version_settings(constraint='>=0.30.0')
load('ext://helm_resource', 'helm_resource', 'helm_repo')
# load docker_build_with_restart
load('ext://restart_process', 'docker_build_with_restart')

# All apps launched by Tilt
apps = {
  'backend': True,
  'frontend': True
}

# backend
if apps['backend']:
  local_resource(
    'backend-swag',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 cd apps/backend/src && swag init',
    deps=['apps/backend/src/controllers', 'apps/backend/src/models', 'apps/backend/src/main.go', 'apps/backend/src/utils'],
    labels=['backend'],
  )
  local_resource(
    'backend-build',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 cd apps/backend/src && go build -o ../build/backend ./ ',
    deps=['apps/backend/src/docs'],
    labels=['backend'],
  )
  docker_build_with_restart(
    'backend-image',
    context='./apps/backend/',
    dockerfile='./apps/backend/Dockerfile',
    entrypoint=['/app/backend'],
    platform='linux/amd64',
    only=['./build',],
    live_update=[
      sync('./apps/backend/build', '/app'),
    ],
  )
  local_resource(
    'backend-helm',
    'helm dependency update apps/backend/helm',
    deps=['apps/backend/helm/Chart.yaml'],
    labels=['backend'],
  )
  helm_resource(
    'backend',
    'apps/backend/helm',
    port_forwards='8081:8080',
    labels=['backend'],
    flags=[
      '-f', 'apps/backend/helm/values.yaml',
    ],
    image_deps=['backend-image'],
    image_keys=[('image', 'tag')],
    deps=['apps/backend/helm/values.yaml'],
  )

# frontend
if apps['frontend']:
  local_resource(
    'frontend-build',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 cd apps/frontend/src && go build -o ../build/frontend ./',
    deps=['apps/frontend/src',],
    labels=['frontend'],
  )
  docker_build_with_restart(
    'frontend-image',
    context='./apps/frontend/',
    dockerfile='./apps/frontend/Dockerfile',
    entrypoint=['/app/frontend'],
    platform='linux/amd64',
    only=['./build', './src/templates',],
    live_update=[
      sync('./apps/frontend/build', '/app'),
      sync('./apps/frontend/src/templates', '/app/templates'),
    ],
  )
  local_resource(
    'frontend-helm',
    'helm dependency update apps/frontend/helm',
    deps=['apps/frontend/helm/Chart.yaml'],
    labels=['frontend'],
  )
  helm_resource(
    'frontend',
    'apps/frontend/helm',
    port_forwards=['8080:8080'],
    labels=['frontend'],
    flags=[
      '-f', 'apps/frontend/helm/values.yaml',
    ],
    image_deps=['frontend-image'],
    image_keys=[('image', 'tag')],
    deps=['apps/frontend/helm/values.yaml'],
  )

# Postgresql
helm_repo('bitnami', 'https://charts.bitnami.com/bitnami', labels=['helm'])
helm_resource(
  'postgresql',
  'bitnami/postgresql',
  port_forwards='5432',
  labels=['postgres'],
  flags=[
    '--version', '^11.0.0',
    '-f', './tilt/postgresql.values.yaml'
  ]
)

# RabbitMQ
k8s_yaml('tilt/rabbitmq.yaml')
k8s_resource(
  'rabbitmq',
  port_forwards=[
      '5672',
      port_forward(15672, 15672, name='Admin (user: guest:guest)')
  ],
  labels=['rabbitmq']
)
k8s_resource(
  new_name='rabbitmq-storage',
  objects=['rabbitmq-pv-volume:persistentvolume', 'rabbitmq-pv-claim:persistentvolumeclaim'],
  labels=['rabbitmq']
)

# Redis
k8s_yaml('tilt/redis.yaml')
k8s_resource(
  'redis',
  labels=['redis'],
  port_forwards='6379',
)
k8s_resource(
  new_name='redis-service',
  objects=['redis-service'],
  labels=['redis']
)
