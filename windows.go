package laago

const msdosHeaderSize = 128
const peHeaderFileHeaderOffset = 4
const peFileHeaderCharacteristicsOffset = 18
const laaOffset = msdosHeaderSize + peHeaderFileHeaderOffset + peFileHeaderCharacteristicsOffset

const peImageFileLargeAddressAware uint8 = 0x0020
