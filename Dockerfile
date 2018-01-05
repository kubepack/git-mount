FROM appscode/kubectl:1.8.0
ADD git-mount /bin/
ENTRYPOINT ["git-mount"]
