case "$1" in
  'start')
    KB=$2
    echo Throttling at $KB Kbytes/s
    sudo ipfw pipe 1 config bw "$KB"KByte/s
    sudo ipfw add 1 pipe 1 src-port 8087
    sudo ipfw add 1 pipe 1 src-port 8088
    sudo ipfw add 1 pipe 1 src-port 8089
    sudo ipfw add 1 pipe 1 src-port 8090
    ;;
  'stop')
    sudo ipfw delete 1
    ;;
esac
