FROM scratch
ADD git-mount /bin/
ENTRYPOINT ["git-mount"]
