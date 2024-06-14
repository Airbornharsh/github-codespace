FROM codercom/code-server:latest

USER root

RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - \
    && apt-get install -y nodejs

RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y \
    && export PATH="$HOME/.cargo/bin:$PATH" \
    && rustup update \
    && rustup default stable

RUN wget -q https://golang.org/dl/go1.18.6.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.18.6.linux-amd64.tar.gz \
    && rm go1.18.6.linux-amd64.tar.gz

RUN apt-get update && apt-get install -y openjdk-17-jdk

RUN apt-get install -y g++ \
    && apt-get install -y build-essential

EXPOSE 8080
EXPOSE 3000
EXPOSE 3001
EXPOSE 3002

# Run CodeServer on startup
CMD ["dumb-init", "code-server", "--bind-addr", "0.0.0.0:8080", "--auth", "none"]
