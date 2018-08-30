#!/bin/bash


# This PLATFORMS list is refreshed after every major Go release.
# Though more platforms may be supported (freebsd/386), they have been removed
# from the standard ports/downloads and therefore removed from this list.
#
PLATFORMS="darwin/amd64" # amd64 only as of go1.5
PLATFORMS="$PLATFORMS linux/amd64"
PLATFORMS="$PLATFORMS linux/ppc64 linux/ppc64le"
PLATFORMS="$PLATFORMS linux/mips64 linux/mips64le"
PLATFORMS="$PLATFORMS freebsd/amd64"
PLATFORMS="$PLATFORMS netbsd/amd64"
PLATFORMS="$PLATFORMS openbsd/amd64"
PLATFORMS="$PLATFORMS dragonfly/amd64"
PLATFORMS="$PLATFORMS solaris/amd64"


PLATFORMS_ARM="linux freebsd netbsd"

##############################################################
# Shouldn't really need to modify anything below this line.  #
##############################################################


VERSION=`geode version`


type setopt >/dev/null 2>&1


make clean
rm -rf release



SCRIPT_NAME=`basename "$0"`
SOURCE_FILE=`echo $@ | sed 's/\.go//'`
CURRENT_DIRECTORY=${PWD##*/}
WORKDIR="./release" # if no src file given, use current dir name


WORKDIRABS=`realpath $WORKDIR`
GODIRABS=`realpath ./pkg/cmd/geode`


echo $WORKDIRABS

echo $GODIRABS

mkdir -p $WORKDIRABS

printf "TARGET           PATH\n"
echo   "========================================"


for PLATFORM in $PLATFORMS; do
  
  cd $WORKDIRABS
  
  GOOS=${PLATFORM%/*}
  GOARCH=${PLATFORM#*/}
  
  NAME="geode-$VERSION-$GOOS-$GOARCH"
  
  TARGETDIR="$WORKDIRABS/$NAME"
  TARGETBINDIR="$TARGETDIR/bin"

  TARNAME="$NAME.tar"

  mkdir -p $TARGETDIR
  mkdir -p $TARGETBINDIR
  
  BIN_FILENAME="$TARGETBINDIR/geode"
  pwd
  
  cp -a "../lib" "$TARGETDIR"
  cp "../Makefile" "$TARGETDIR"
  cp "../stdlib.mk" "$TARGETDIR"
  
  # Special case for windows - add .exe
  if [[ "${GOOS}" == "windows" ]]; then
    BIN_FILENAME="${BIN_FILENAME}.exe";
  fi
  
  
  
  cd $GODIRABS
  CMD="GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${BIN_FILENAME} $@"
  echo $CMD
  eval $CMD || exit 1
  cd $WORKDIRABS
  
  
  tar -cf $TARNAME $NAME
  rm -rf $TARGETDIR
  
  # printf "%.20s %22s\n" "${GOOS}-${GOARCH}" "`realpath $TARNAME`"
  
done

# rm -rf working
