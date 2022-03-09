#!/usr/bin/env bash
# shellcheck disable=SC2268

rm -rf /usr/local/gredir
rm -f /etc/systemd/system/gredir.service
mkdir /usr/local/gredir
cp * /usr/local/gredir

chmod u+x /usr/local/gredir/gredir

cat > /etc/systemd/system/gredir.service <<EOF
[Unit]
Description=gredir tcp redirector

[Service]
Type=notify
ExecStart=/usr/local/gredir/gredir -c /usr/local/gredir/config.yaml
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF

systemctl daemon-reload
systemctl start gredir
systemctl enable gredir