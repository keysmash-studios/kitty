#!/bin/sh

[ "$(id -u)" -gt 0 ] && echo "Run with more privileges!" && exit 1

for i in "/usr/bin" "/usr/local/bin" "/bin"; do
	DIR="$i"; break
done

case "$(uname)" in
	*Linux*) BIN="linux";;
	*Darwin*) BIN="macos";;
	*) echo "Unknown OS"; exit 1 ;;
esac

[ ! -f "kitty-$BIN" ] && echo "Can't find binary! (kitty-$BIN)" && exit 1
cp "kitty-$BIN" "$DIR"/kitty || {
	echo "Can't move the binary..."
	echo "Does it exist? (kitty-$BIN)"
	exit 1
}
chmod 755 "$DIR"/kitty
