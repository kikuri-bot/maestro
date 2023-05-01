# Reset just in case
if [ "$1" != "-reset"  ] || [ "$1" != "-r"  ]; then
    echo "Clearing old configuration..."
    systemctl stop maestro.service
    rm "/lib/systemd/system/maestro.service"
    systemctl daemon-reload
    journalctl --rotate
    journalctl --vacuum-time=1s
    chmod 777 "/etc/kikuri" -R
    rm /etc/kikuri/maestro/service/maestro.o
fi

# Configure systemctl daemon
echo "Compiling..."
/usr/local/go/bin/go build -o /etc/kikuri/maestro/service/maestro.o ./main.go
chmod 777 "/etc/kikuri" -R

echo "Registering new daemon..."
cp "/etc/kikuri/maestro/service/maestro.service" "/lib/systemd/system/maestro.service"
chmod 644 "/lib/systemd/system/maestro.service"
systemctl daemon-reload
systemctl enable maestro.service
systemctl restart maestro
echo "Finished!"