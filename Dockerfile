FROM centos
COPY ./nexus-cli    /nexus/nexus-cli
COPY ./clean-up-nexus-repo.sh /nexus/clean-up-nexus-repo.sh
CMD ["/nexus/clean-up-nexus-repo.sh"]

