# version_settings() enforces a minimum Tilt version
# https://docs.tilt.dev/api.html#api.version_settings
version_settings(constraint='>=0.30.0')
load('ext://helm_resource', 'helm_resource', 'helm_repo')
load('ext://restart_process', 'docker_build_with_restart')

# All apps launched by Tilt
apps = {
  'backend': True,
  'frontend': True
}

# backend
if apps['backend']:
  local_resource(
    'backend-compilation',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 cd apps/backend/src && go build -o ../build/backend ./',
    deps=['apps/backend/src'],
    labels=['backend'],
  )
  docker_build_with_restart(
    'backend',
    context='./apps/backend/',
    dockerfile='./apps/backend/Dockerfile',
    entrypoint=['/app/backend'],
    only=['./build'],
    platform='linux/amd64',
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
    image_deps=['backend'],
    image_keys=[('image', 'tag')],
  )

# frontend
if apps['frontend']:
  local_resource(
    'frontend-compilation',
    'CGO_ENABLED=0 GOOS=linux GOARCH=amd64 cd apps/frontend && go build -o build/frontend ./',
    deps=['apps/frontend/main.go', 'apps/frontend/constants.go',],
    resource_deps = ['deploy'],
    labels=['frontend'],
  )
  docker_build_with_restart(
    'frontend-image',
    context='./apps/frontend/',
    entrypoint=['/app/frontend'],
    dockerfile='./apps/frontend/Dockerfile',
    platform='linux/amd64',
    only=['./build', './templates',],
    live_update=[
      sync('./apps/frontend/build', '/app'),
      sync('./apps/frontend/templates', '/app/templates'),
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
