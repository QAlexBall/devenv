# Use the latest Golang image with version 1.23.3 as the base image
FROM golang:1.23.3-bullseye

# Set environment variables for Go modules and Kubebuilder
ENV GO111MODULE=on \
    CGO_ENABLED=0

# Install dependencies
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    make \
    bash \
    curl \
    unzip \
    net-tools netcat-openbsd curl wget iputils-ping \
    docker.io && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Install Kubebuilder
ENV KUBEBUILDER_VERSION=4.3.1
RUN curl -L https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${KUBEBUILDER_VERSION}/kubebuilder_linux_amd64 \
    -o /usr/local/kubebuilder && \
    chmod +x /usr/local/kubebuilder && \
    mv /usr/local/kubebuilder /usr/local/bin/kubebuilder

ENV KUBECTL_VERSION=v1.29.3
RUN curl -LO "https://dl.k8s.io/release/${KUBECTL_VERSION}/bin/linux/amd64/kubectl" && \
    mv kubectl /usr/local/bin/kubectl && \
    chmod +x /usr/local/bin/kubectl

# Set up workspace
WORKDIR /workspace
ENV KUBECONFIG=/workspace/.kube/config

# Expose application port (optional, adjust as needed)
EXPOSE 8080

# Set the default command (bash for interactive development)
CMD ["bash"]
