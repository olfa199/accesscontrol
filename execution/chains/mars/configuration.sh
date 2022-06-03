crossdomaind  tx crossdomain configure-local-domain   "" "" "" ""   --from alice --yes --home ~/.mars --chain-id mars --node tcp://localhost:26659 --gas=auto --gas-adjustment=1.15
crossdomaind query crossdomain show-local-domain --node tcp://localhost:26659
crossdomaind query crossdomain show-local-domain-certificate --node tcp://localhost:26659
crossdomaind query crossdomain  show-private-key  --node tcp://localhost:26659 
crossdomaind query crossdomain show-root-certificate --node tcp://localhost:26659

crossdomaind tx crossdomain  create-forward-policy "broadcast" "" "" --from alice --yes --home ~/.mars --chain-id mars --node tcp://localhost:26659 --gas=auto --gas-adjustment=1.15
crossdomaind query crossdomain show-forward-policy --node tcp://localhost:26659

crossdomaind tx crossdomain  update-decision-policy 0 0 '' '' '{"notBefore":"2026-01-01 00:00:00","notAfter":"2026-01-01 00:00:00"}' "" --from alice --yes --home ~/.mars --chain-id mars --node tcp://localhost:26659 --gas=auto --gas-adjustment=1.15
crossdomaind query crossdomain show-decision-policy --node tcp://localhost:26659