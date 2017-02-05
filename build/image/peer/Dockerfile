FROM hyperledger/fabric-runtime:x86_64-0.7.0-snapshot-00ce328
ENV PEER_CFG_PATH /etc/hyperledger/fabric
RUN mkdir -p /var/hyperledger/db $PEER_CFG_PATH
COPY payload/peer /usr/local/bin
COPY payload/core.yaml $PEER_CFG_PATH
ADD  payload/msp-sampleconfig.tar.bz2 $PEER_CFG_PATH
ADD  payload/genesis-sampleconfig.tar.bz2 $PEER_CFG_PATH
CMD peer node start
