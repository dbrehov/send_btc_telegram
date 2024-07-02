go mod tidy
go mod init github.com/dbrehov/send_btc_telegram
git add . &&  git commit -m "ok" && git push


export CGO_CFLAGS="-I/data/data/com.termux/files/usr/include"
export CGO_LDFLAGS="-L/data/data/com.termux/files/usr/lib"


cp -r android-ndk-r21e/sysroot/usr/include/* /data/data/com.termux/files/usr/include/

ls /data/data/com.termux/files/usr/include/stdint.h

cp -r android-ndk-r21e/sysroot/usr/include/* /data/data/com.termux/files/usr/include/

wget https://dl.google.com/android/repository/android-ndk-r21e-linux-x86_64.zip

unzip android-ndk-r21e-linux-x86_64.zip

find android-ndk-r21e/ -name "stdarg.h"

ls android-ndk-r21e/sysroot/usr/include/

cp android-ndk-r21e/sysroot/usr/include/stdarg.h /data/data/com.termux/files/usr/include/

export CGO_CFLAGS="-I/path/to/stdarg_directory"
export CGO_LDFLAGS="-L/path/to/library_directory"

ls /data/data/com.termux/files/usr/include/


cp -r android-ndk-r21e/sysroot/usr/include/* /data/data/com.termux/files/usr/include/

pkg install clang

find /data/data/com.termux/files/usr/include -name "stdarg.h"

find /data/data/com.termux/files/usr/include -name "stdlib.h"

pkg install clang make

pkg install --reinstall clang

pkg install clang

pkg update

pkg upgrade

pkg install openssh

passwd



