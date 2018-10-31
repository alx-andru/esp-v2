# Copyright 2018 Google Cloud Platform Proxy Authors
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM gcc:8
LABEL maintainer="jilinxia@google.com"

# add env we can debug with the image name:tag
ARG IMAGE_ARG
ENV IMAGE=${IMAGE_ARG}

RUN apt-get update && \
    apt-get -y install \
    wget make cmake python python-pip pkg-config realpath \
    zlib1g-dev curl libtool automake zip time rsync ninja-build \
    git bash-completion

# clang is used for TSAN and ASAN tests
RUN sh -c 'curl http://apt.llvm.org/llvm-snapshot.gpg.key | apt-key add -'
RUN sh -c 'echo "deb http://apt.llvm.org/stretch/ llvm-toolchain-stretch-6.0 main" > /etc/apt/sources.list.d/llvm.list'

# install bazel
RUN INSTALLER="bazel-0.18.0-installer-linux-x86_64.sh"; \
    DOWNLOAD_URL="https://github.com/bazelbuild/bazel/releases/download/0.18.0/${INSTALLER}"; \
    wget -q "${DOWNLOAD_URL}" && \
    chmod +x "${INSTALLER}" && "./${INSTALLER}" && rm "${INSTALLER}"

# install golang and setup Go standard envs
ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH
ENV PATH $GOPATH/bin:$PATH

ENV GO_TARBALL "go1.11.1.linux-amd64.tar.gz"
RUN wget -q "https://storage.googleapis.com/golang/${GO_TARBALL}" && \
    tar xzf "${GO_TARBALL}" -C /usr/local && \
    rm "${GO_TARBALL}"
