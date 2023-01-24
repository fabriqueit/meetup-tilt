# meetup-tilt

## Prerequisite

- Docker/containerd
- Tilt
- kubernetes local
  - k3d
    curl -s https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
    k3d cluster create meetup-tilt --registry-create meetup-tilt-registry
