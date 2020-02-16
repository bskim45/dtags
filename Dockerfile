FROM gcr.io/distroless/static:nonroot
LABEL maintainer="bskim45@gmail.com"

ARG VCS_REF
ARG BUILD_DATE
ARG VERSION

LABEL org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.name="dtags" \
    org.label-schema.url="https://hub.docker.com/r/bskim45/dtags/" \
    org.label-schema.vcs-url="https://github.com/bskim45/dtags" \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.version=$VERSION

COPY dist/dtags_linux_amd64/dtags /go/bin/dtags

ENTRYPOINT ["/go/bin/dtags"]
