package AES

import (
	"encoding/binary"
	"fmt"
)

var g_Box1 = [...]uint8{
	0x63, 0x7C, 0x77, 0x7B, 0xF2, 0x6B, 0x6F, 0xC5, 0x30, 0x01, 0x67, 0x2B, 0xFE, 0xD7, 0xAB, 0x76,
	0xCA, 0x82, 0xC9, 0x7D, 0xFA, 0x59, 0x47, 0xF0, 0xAD, 0xD4, 0xA2, 0xAF, 0x9C, 0xA4, 0x72, 0xC0,
	0xB7, 0xFD, 0x93, 0x26, 0x36, 0x3F, 0xF7, 0xCC, 0x34, 0xA5, 0xE5, 0xF1, 0x71, 0xD8, 0x31, 0x15,
	0x04, 0xC7, 0x23, 0xC3, 0x18, 0x96, 0x05, 0x9A, 0x07, 0x12, 0x80, 0xE2, 0xEB, 0x27, 0xB2, 0x75,
	0x09, 0x83, 0x2C, 0x1A, 0x1B, 0x6E, 0x5A, 0xA0, 0x52, 0x3B, 0xD6, 0xB3, 0x29, 0xE3, 0x2F, 0x84,
	0x53, 0xD1, 0x00, 0xED, 0x20, 0xFC, 0xB1, 0x5B, 0x6A, 0xCB, 0xBE, 0x39, 0x4A, 0x4C, 0x58, 0xCF,
	0xD0, 0xEF, 0xAA, 0xFB, 0x43, 0x4D, 0x33, 0x85, 0x45, 0xF9, 0x02, 0x7F, 0x50, 0x3C, 0x9F, 0xA8,
	0x51, 0xA3, 0x40, 0x8F, 0x92, 0x9D, 0x38, 0xF5, 0xBC, 0xB6, 0xDA, 0x21, 0x10, 0xFF, 0xF3, 0xD2,
	0xCD, 0x0C, 0x13, 0xEC, 0x5F, 0x97, 0x44, 0x17, 0xC4, 0xA7, 0x7E, 0x3D, 0x64, 0x5D, 0x19, 0x73,
	0x60, 0x81, 0x4F, 0xDC, 0x22, 0x2A, 0x90, 0x88, 0x46, 0xEE, 0xB8, 0x14, 0xDE, 0x5E, 0x0B, 0xDB,
	0xE0, 0x32, 0x3A, 0x0A, 0x49, 0x06, 0x24, 0x5C, 0xC2, 0xD3, 0xAC, 0x62, 0x91, 0x95, 0xE4, 0x79,
	0xE7, 0xC8, 0x37, 0x6D, 0x8D, 0xD5, 0x4E, 0xA9, 0x6C, 0x56, 0xF4, 0xEA, 0x65, 0x7A, 0xAE, 0x08,
	0xBA, 0x78, 0x25, 0x2E, 0x1C, 0xA6, 0xB4, 0xC6, 0xE8, 0xDD, 0x74, 0x1F, 0x4B, 0xBD, 0x8B, 0x8A,
	0x70, 0x3E, 0xB5, 0x66, 0x48, 0x03, 0xF6, 0x0E, 0x61, 0x35, 0x57, 0xB9, 0x86, 0xC1, 0x1D, 0x9E,
	0xE1, 0xF8, 0x98, 0x11, 0x69, 0xD9, 0x8E, 0x94, 0x9B, 0x1E, 0x87, 0xE9, 0xCE, 0x55, 0x28, 0xDF,
	0x8C, 0xA1, 0x89, 0x0D, 0xBF, 0xE6, 0x42, 0x68, 0x41, 0x99, 0x2D, 0x0F, 0xB0, 0x54, 0xBB, 0x16,
}
var g_Box2 = [...]uint8{
	0x8D, 0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1B, 0x36, 0x6C, 0xD8, 0xAB, 0x4D, 0x9A,
	0x2F, 0x5E, 0xBC, 0x63, 0xC6, 0x97, 0x35, 0x6A, 0xD4, 0xB3, 0x7D, 0xFA, 0xEF, 0xC5, 0x91, 0x39,
	0x72, 0xE4, 0xD3, 0xBD, 0x61, 0xC2, 0x9F, 0x25, 0x4A, 0x94, 0x33, 0x66, 0xCC, 0x83, 0x1D, 0x3A,
	0x74, 0xE8, 0xCB, 0x8D, 0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1B, 0x36, 0x6C, 0xD8,
	0xAB, 0x4D, 0x9A, 0x2F, 0x5E, 0xBC, 0x63, 0xC6, 0x97, 0x35, 0x6A, 0xD4, 0xB3, 0x7D, 0xFA, 0xEF,
	0xC5, 0x91, 0x39, 0x72, 0xE4, 0xD3, 0xBD, 0x61, 0xC2, 0x9F, 0x25, 0x4A, 0x94, 0x33, 0x66, 0xCC,
	0x83, 0x1D, 0x3A, 0x74, 0xE8, 0xCB, 0x8D, 0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80, 0x1B,
	0x36, 0x6C, 0xD8, 0xAB, 0x4D, 0x9A, 0x2F, 0x5E, 0xBC, 0x63, 0xC6, 0x97, 0x35, 0x6A, 0xD4, 0xB3,
	0x7D, 0xFA, 0xEF, 0xC5, 0x91, 0x39, 0x72, 0xE4, 0xD3, 0xBD, 0x61, 0xC2, 0x9F, 0x25, 0x4A, 0x94,
	0x33, 0x66, 0xCC, 0x83, 0x1D, 0x3A, 0x74, 0xE8, 0xCB, 0x8D, 0x01, 0x02, 0x04, 0x08, 0x10, 0x20,
	0x40, 0x80, 0x1B, 0x36, 0x6C, 0xD8, 0xAB, 0x4D, 0x9A, 0x2F, 0x5E, 0xBC, 0x63, 0xC6, 0x97, 0x35,
	0x6A, 0xD4, 0xB3, 0x7D, 0xFA, 0xEF, 0xC5, 0x91, 0x39, 0x72, 0xE4, 0xD3, 0xBD, 0x61, 0xC2, 0x9F,
	0x25, 0x4A, 0x94, 0x33, 0x66, 0xCC, 0x83, 0x1D, 0x3A, 0x74, 0xE8, 0xCB, 0x8D, 0x01, 0x02, 0x04,
	0x08, 0x10, 0x20, 0x40, 0x80, 0x1B, 0x36, 0x6C, 0xD8, 0xAB, 0x4D, 0x9A, 0x2F, 0x5E, 0xBC, 0x63,
	0xC6, 0x97, 0x35, 0x6A, 0xD4, 0xB3, 0x7D, 0xFA, 0xEF, 0xC5, 0x91, 0x39, 0x72, 0xE4, 0xD3, 0xBD,
	0x61, 0xC2, 0x9F, 0x25, 0x4A, 0x94, 0x33, 0x66, 0xCC, 0x83, 0x1D, 0x3A, 0x74, 0xE8, 0xCB, 0x8D,
}
var g_Box3 = [...]uint8{
	0x00, 0x02, 0x04, 0x06, 0x08, 0x0A, 0x0C, 0x0E, 0x10, 0x12, 0x14, 0x16, 0x18, 0x1A, 0x1C, 0x1E,
	0x20, 0x22, 0x24, 0x26, 0x28, 0x2A, 0x2C, 0x2E, 0x30, 0x32, 0x34, 0x36, 0x38, 0x3A, 0x3C, 0x3E,
	0x40, 0x42, 0x44, 0x46, 0x48, 0x4A, 0x4C, 0x4E, 0x50, 0x52, 0x54, 0x56, 0x58, 0x5A, 0x5C, 0x5E,
	0x60, 0x62, 0x64, 0x66, 0x68, 0x6A, 0x6C, 0x6E, 0x70, 0x72, 0x74, 0x76, 0x78, 0x7A, 0x7C, 0x7E,
	0x80, 0x82, 0x84, 0x86, 0x88, 0x8A, 0x8C, 0x8E, 0x90, 0x92, 0x94, 0x96, 0x98, 0x9A, 0x9C, 0x9E,
	0xA0, 0xA2, 0xA4, 0xA6, 0xA8, 0xAA, 0xAC, 0xAE, 0xB0, 0xB2, 0xB4, 0xB6, 0xB8, 0xBA, 0xBC, 0xBE,
	0xC0, 0xC2, 0xC4, 0xC6, 0xC8, 0xCA, 0xCC, 0xCE, 0xD0, 0xD2, 0xD4, 0xD6, 0xD8, 0xDA, 0xDC, 0xDE,
	0xE0, 0xE2, 0xE4, 0xE6, 0xE8, 0xEA, 0xEC, 0xEE, 0xF0, 0xF2, 0xF4, 0xF6, 0xF8, 0xFA, 0xFC, 0xFE,
	0x1B, 0x19, 0x1F, 0x1D, 0x13, 0x11, 0x17, 0x15, 0x0B, 0x09, 0x0F, 0x0D, 0x03, 0x01, 0x07, 0x05,
	0x3B, 0x39, 0x3F, 0x3D, 0x33, 0x31, 0x37, 0x35, 0x2B, 0x29, 0x2F, 0x2D, 0x23, 0x21, 0x27, 0x25,
	0x5B, 0x59, 0x5F, 0x5D, 0x53, 0x51, 0x57, 0x55, 0x4B, 0x49, 0x4F, 0x4D, 0x43, 0x41, 0x47, 0x45,
	0x7B, 0x79, 0x7F, 0x7D, 0x73, 0x71, 0x77, 0x75, 0x6B, 0x69, 0x6F, 0x6D, 0x63, 0x61, 0x67, 0x65,
	0x9B, 0x99, 0x9F, 0x9D, 0x93, 0x91, 0x97, 0x95, 0x8B, 0x89, 0x8F, 0x8D, 0x83, 0x81, 0x87, 0x85,
	0xBB, 0xB9, 0xBF, 0xBD, 0xB3, 0xB1, 0xB7, 0xB5, 0xAB, 0xA9, 0xAF, 0xAD, 0xA3, 0xA1, 0xA7, 0xA5,
	0xDB, 0xD9, 0xDF, 0xDD, 0xD3, 0xD1, 0xD7, 0xD5, 0xCB, 0xC9, 0xCF, 0xCD, 0xC3, 0xC1, 0xC7, 0xC5,
	0xFB, 0xF9, 0xFF, 0xFD, 0xF3, 0xF1, 0xF7, 0xF5, 0xEB, 0xE9, 0xEF, 0xED, 0xE3, 0xE1, 0xE7, 0xE5,
}
var g_Box4 = [...]uint8{
	0x00, 0x03, 0x06, 0x05, 0x0C, 0x0F, 0x0A, 0x09, 0x18, 0x1B, 0x1E, 0x1D, 0x14, 0x17, 0x12, 0x11,
	0x30, 0x33, 0x36, 0x35, 0x3C, 0x3F, 0x3A, 0x39, 0x28, 0x2B, 0x2E, 0x2D, 0x24, 0x27, 0x22, 0x21,
	0x60, 0x63, 0x66, 0x65, 0x6C, 0x6F, 0x6A, 0x69, 0x78, 0x7B, 0x7E, 0x7D, 0x74, 0x77, 0x72, 0x71,
	0x50, 0x53, 0x56, 0x55, 0x5C, 0x5F, 0x5A, 0x59, 0x48, 0x4B, 0x4E, 0x4D, 0x44, 0x47, 0x42, 0x41,
	0xC0, 0xC3, 0xC6, 0xC5, 0xCC, 0xCF, 0xCA, 0xC9, 0xD8, 0xDB, 0xDE, 0xDD, 0xD4, 0xD7, 0xD2, 0xD1,
	0xF0, 0xF3, 0xF6, 0xF5, 0xFC, 0xFF, 0xFA, 0xF9, 0xE8, 0xEB, 0xEE, 0xED, 0xE4, 0xE7, 0xE2, 0xE1,
	0xA0, 0xA3, 0xA6, 0xA5, 0xAC, 0xAF, 0xAA, 0xA9, 0xB8, 0xBB, 0xBE, 0xBD, 0xB4, 0xB7, 0xB2, 0xB1,
	0x90, 0x93, 0x96, 0x95, 0x9C, 0x9F, 0x9A, 0x99, 0x88, 0x8B, 0x8E, 0x8D, 0x84, 0x87, 0x82, 0x81,
	0x9B, 0x98, 0x9D, 0x9E, 0x97, 0x94, 0x91, 0x92, 0x83, 0x80, 0x85, 0x86, 0x8F, 0x8C, 0x89, 0x8A,
	0xAB, 0xA8, 0xAD, 0xAE, 0xA7, 0xA4, 0xA1, 0xA2, 0xB3, 0xB0, 0xB5, 0xB6, 0xBF, 0xBC, 0xB9, 0xBA,
	0xFB, 0xF8, 0xFD, 0xFE, 0xF7, 0xF4, 0xF1, 0xF2, 0xE3, 0xE0, 0xE5, 0xE6, 0xEF, 0xEC, 0xE9, 0xEA,
	0xCB, 0xC8, 0xCD, 0xCE, 0xC7, 0xC4, 0xC1, 0xC2, 0xD3, 0xD0, 0xD5, 0xD6, 0xDF, 0xDC, 0xD9, 0xDA,
	0x5B, 0x58, 0x5D, 0x5E, 0x57, 0x54, 0x51, 0x52, 0x43, 0x40, 0x45, 0x46, 0x4F, 0x4C, 0x49, 0x4A,
	0x6B, 0x68, 0x6D, 0x6E, 0x67, 0x64, 0x61, 0x62, 0x73, 0x70, 0x75, 0x76, 0x7F, 0x7C, 0x79, 0x7A,
	0x3B, 0x38, 0x3D, 0x3E, 0x37, 0x34, 0x31, 0x32, 0x23, 0x20, 0x25, 0x26, 0x2F, 0x2C, 0x29, 0x2A,
	0x0B, 0x08, 0x0D, 0x0E, 0x07, 0x04, 0x01, 0x02, 0x13, 0x10, 0x15, 0x16, 0x1F, 0x1C, 0x19, 0x1A,
}
var g_DereplaceFourKeyBox = [...]uint8{
	0x52, 0x09, 0x6A, 0xD5, 0x30, 0x36, 0xA5, 0x38, 0xBF, 0x40, 0xA3, 0x9E, 0x81, 0xF3, 0xD7, 0xFB,
	0x7C, 0xE3, 0x39, 0x82, 0x9B, 0x2F, 0xFF, 0x87, 0x34, 0x8E, 0x43, 0x44, 0xC4, 0xDE, 0xE9, 0xCB,
	0x54, 0x7B, 0x94, 0x32, 0xA6, 0xC2, 0x23, 0x3D, 0xEE, 0x4C, 0x95, 0x0B, 0x42, 0xFA, 0xC3, 0x4E,
	0x08, 0x2E, 0xA1, 0x66, 0x28, 0xD9, 0x24, 0xB2, 0x76, 0x5B, 0xA2, 0x49, 0x6D, 0x8B, 0xD1, 0x25,
	0x72, 0xF8, 0xF6, 0x64, 0x86, 0x68, 0x98, 0x16, 0xD4, 0xA4, 0x5C, 0xCC, 0x5D, 0x65, 0xB6, 0x92,
	0x6C, 0x70, 0x48, 0x50, 0xFD, 0xED, 0xB9, 0xDA, 0x5E, 0x15, 0x46, 0x57, 0xA7, 0x8D, 0x9D, 0x84,
	0x90, 0xD8, 0xAB, 0x00, 0x8C, 0xBC, 0xD3, 0x0A, 0xF7, 0xE4, 0x58, 0x05, 0xB8, 0xB3, 0x45, 0x06,
	0xD0, 0x2C, 0x1E, 0x8F, 0xCA, 0x3F, 0x0F, 0x02, 0xC1, 0xAF, 0xBD, 0x03, 0x01, 0x13, 0x8A, 0x6B,
	0x3A, 0x91, 0x11, 0x41, 0x4F, 0x67, 0xDC, 0xEA, 0x97, 0xF2, 0xCF, 0xCE, 0xF0, 0xB4, 0xE6, 0x73,
	0x96, 0xAC, 0x74, 0x22, 0xE7, 0xAD, 0x35, 0x85, 0xE2, 0xF9, 0x37, 0xE8, 0x1C, 0x75, 0xDF, 0x6E,
	0x47, 0xF1, 0x1A, 0x71, 0x1D, 0x29, 0xC5, 0x89, 0x6F, 0xB7, 0x62, 0x0E, 0xAA, 0x18, 0xBE, 0x1B,
	0xFC, 0x56, 0x3E, 0x4B, 0xC6, 0xD2, 0x79, 0x20, 0x9A, 0xDB, 0xC0, 0xFE, 0x78, 0xCD, 0x5A, 0xF4,
	0x1F, 0xDD, 0xA8, 0x33, 0x88, 0x07, 0xC7, 0x31, 0xB1, 0x12, 0x10, 0x59, 0x27, 0x80, 0xEC, 0x5F,
	0x60, 0x51, 0x7F, 0xA9, 0x19, 0xB5, 0x4A, 0x0D, 0x2D, 0xE5, 0x7A, 0x9F, 0x93, 0xC9, 0x9C, 0xEF,
	0xA0, 0xE0, 0x3B, 0x4D, 0xAE, 0x2A, 0xF5, 0xB0, 0xC8, 0xEB, 0xBB, 0x3C, 0x83, 0x53, 0x99, 0x61,
	0x17, 0x2B, 0x04, 0x7E, 0xBA, 0x77, 0xD6, 0x26, 0xE1, 0x69, 0x14, 0x63, 0x55, 0x21, 0x0C, 0x7D,
}
var g_DeBox1 = [...]uint8{
	0x00, 0x0E, 0x1C, 0x12, 0x38, 0x36, 0x24, 0x2A, 0x70, 0x7E, 0x6C, 0x62, 0x48, 0x46, 0x54, 0x5A,
	0xE0, 0xEE, 0xFC, 0xF2, 0xD8, 0xD6, 0xC4, 0xCA, 0x90, 0x9E, 0x8C, 0x82, 0xA8, 0xA6, 0xB4, 0xBA,
	0xDB, 0xD5, 0xC7, 0xC9, 0xE3, 0xED, 0xFF, 0xF1, 0xAB, 0xA5, 0xB7, 0xB9, 0x93, 0x9D, 0x8F, 0x81,
	0x3B, 0x35, 0x27, 0x29, 0x03, 0x0D, 0x1F, 0x11, 0x4B, 0x45, 0x57, 0x59, 0x73, 0x7D, 0x6F, 0x61,
	0xAD, 0xA3, 0xB1, 0xBF, 0x95, 0x9B, 0x89, 0x87, 0xDD, 0xD3, 0xC1, 0xCF, 0xE5, 0xEB, 0xF9, 0xF7,
	0x4D, 0x43, 0x51, 0x5F, 0x75, 0x7B, 0x69, 0x67, 0x3D, 0x33, 0x21, 0x2F, 0x05, 0x0B, 0x19, 0x17,
	0x76, 0x78, 0x6A, 0x64, 0x4E, 0x40, 0x52, 0x5C, 0x06, 0x08, 0x1A, 0x14, 0x3E, 0x30, 0x22, 0x2C,
	0x96, 0x98, 0x8A, 0x84, 0xAE, 0xA0, 0xB2, 0xBC, 0xE6, 0xE8, 0xFA, 0xF4, 0xDE, 0xD0, 0xC2, 0xCC,
	0x41, 0x4F, 0x5D, 0x53, 0x79, 0x77, 0x65, 0x6B, 0x31, 0x3F, 0x2D, 0x23, 0x09, 0x07, 0x15, 0x1B,
	0xA1, 0xAF, 0xBD, 0xB3, 0x99, 0x97, 0x85, 0x8B, 0xD1, 0xDF, 0xCD, 0xC3, 0xE9, 0xE7, 0xF5, 0xFB,
	0x9A, 0x94, 0x86, 0x88, 0xA2, 0xAC, 0xBE, 0xB0, 0xEA, 0xE4, 0xF6, 0xF8, 0xD2, 0xDC, 0xCE, 0xC0,
	0x7A, 0x74, 0x66, 0x68, 0x42, 0x4C, 0x5E, 0x50, 0x0A, 0x04, 0x16, 0x18, 0x32, 0x3C, 0x2E, 0x20,
	0xEC, 0xE2, 0xF0, 0xFE, 0xD4, 0xDA, 0xC8, 0xC6, 0x9C, 0x92, 0x80, 0x8E, 0xA4, 0xAA, 0xB8, 0xB6,
	0x0C, 0x02, 0x10, 0x1E, 0x34, 0x3A, 0x28, 0x26, 0x7C, 0x72, 0x60, 0x6E, 0x44, 0x4A, 0x58, 0x56,
	0x37, 0x39, 0x2B, 0x25, 0x0F, 0x01, 0x13, 0x1D, 0x47, 0x49, 0x5B, 0x55, 0x7F, 0x71, 0x63, 0x6D,
	0xD7, 0xD9, 0xCB, 0xC5, 0xEF, 0xE1, 0xF3, 0xFD, 0xA7, 0xA9, 0xBB, 0xB5, 0x9F, 0x91, 0x83, 0x8D,
}
var g_DeBox2 = [...]uint8{
	0x00, 0x0D, 0x1A, 0x17, 0x34, 0x39, 0x2E, 0x23, 0x68, 0x65, 0x72, 0x7F, 0x5C, 0x51, 0x46, 0x4B,
	0xD0, 0xDD, 0xCA, 0xC7, 0xE4, 0xE9, 0xFE, 0xF3, 0xB8, 0xB5, 0xA2, 0xAF, 0x8C, 0x81, 0x96, 0x9B,
	0xBB, 0xB6, 0xA1, 0xAC, 0x8F, 0x82, 0x95, 0x98, 0xD3, 0xDE, 0xC9, 0xC4, 0xE7, 0xEA, 0xFD, 0xF0,
	0x6B, 0x66, 0x71, 0x7C, 0x5F, 0x52, 0x45, 0x48, 0x03, 0x0E, 0x19, 0x14, 0x37, 0x3A, 0x2D, 0x20,
	0x6D, 0x60, 0x77, 0x7A, 0x59, 0x54, 0x43, 0x4E, 0x05, 0x08, 0x1F, 0x12, 0x31, 0x3C, 0x2B, 0x26,
	0xBD, 0xB0, 0xA7, 0xAA, 0x89, 0x84, 0x93, 0x9E, 0xD5, 0xD8, 0xCF, 0xC2, 0xE1, 0xEC, 0xFB, 0xF6,
	0xD6, 0xDB, 0xCC, 0xC1, 0xE2, 0xEF, 0xF8, 0xF5, 0xBE, 0xB3, 0xA4, 0xA9, 0x8A, 0x87, 0x90, 0x9D,
	0x06, 0x0B, 0x1C, 0x11, 0x32, 0x3F, 0x28, 0x25, 0x6E, 0x63, 0x74, 0x79, 0x5A, 0x57, 0x40, 0x4D,
	0xDA, 0xD7, 0xC0, 0xCD, 0xEE, 0xE3, 0xF4, 0xF9, 0xB2, 0xBF, 0xA8, 0xA5, 0x86, 0x8B, 0x9C, 0x91,
	0x0A, 0x07, 0x10, 0x1D, 0x3E, 0x33, 0x24, 0x29, 0x62, 0x6F, 0x78, 0x75, 0x56, 0x5B, 0x4C, 0x41,
	0x61, 0x6C, 0x7B, 0x76, 0x55, 0x58, 0x4F, 0x42, 0x09, 0x04, 0x13, 0x1E, 0x3D, 0x30, 0x27, 0x2A,
	0xB1, 0xBC, 0xAB, 0xA6, 0x85, 0x88, 0x9F, 0x92, 0xD9, 0xD4, 0xC3, 0xCE, 0xED, 0xE0, 0xF7, 0xFA,
	0xB7, 0xBA, 0xAD, 0xA0, 0x83, 0x8E, 0x99, 0x94, 0xDF, 0xD2, 0xC5, 0xC8, 0xEB, 0xE6, 0xF1, 0xFC,
	0x67, 0x6A, 0x7D, 0x70, 0x53, 0x5E, 0x49, 0x44, 0x0F, 0x02, 0x15, 0x18, 0x3B, 0x36, 0x21, 0x2C,
	0x0C, 0x01, 0x16, 0x1B, 0x38, 0x35, 0x22, 0x2F, 0x64, 0x69, 0x7E, 0x73, 0x50, 0x5D, 0x4A, 0x47,
	0xDC, 0xD1, 0xC6, 0xCB, 0xE8, 0xE5, 0xF2, 0xFF, 0xB4, 0xB9, 0xAE, 0xA3, 0x80, 0x8D, 0x9A, 0x97,
}
var g_DeBox3 = [...]uint8{
	0x00, 0x09, 0x12, 0x1B, 0x24, 0x2D, 0x36, 0x3F, 0x48, 0x41, 0x5A, 0x53, 0x6C, 0x65, 0x7E, 0x77,
	0x90, 0x99, 0x82, 0x8B, 0xB4, 0xBD, 0xA6, 0xAF, 0xD8, 0xD1, 0xCA, 0xC3, 0xFC, 0xF5, 0xEE, 0xE7,
	0x3B, 0x32, 0x29, 0x20, 0x1F, 0x16, 0x0D, 0x04, 0x73, 0x7A, 0x61, 0x68, 0x57, 0x5E, 0x45, 0x4C,
	0xAB, 0xA2, 0xB9, 0xB0, 0x8F, 0x86, 0x9D, 0x94, 0xE3, 0xEA, 0xF1, 0xF8, 0xC7, 0xCE, 0xD5, 0xDC,
	0x76, 0x7F, 0x64, 0x6D, 0x52, 0x5B, 0x40, 0x49, 0x3E, 0x37, 0x2C, 0x25, 0x1A, 0x13, 0x08, 0x01,
	0xE6, 0xEF, 0xF4, 0xFD, 0xC2, 0xCB, 0xD0, 0xD9, 0xAE, 0xA7, 0xBC, 0xB5, 0x8A, 0x83, 0x98, 0x91,
	0x4D, 0x44, 0x5F, 0x56, 0x69, 0x60, 0x7B, 0x72, 0x05, 0x0C, 0x17, 0x1E, 0x21, 0x28, 0x33, 0x3A,
	0xDD, 0xD4, 0xCF, 0xC6, 0xF9, 0xF0, 0xEB, 0xE2, 0x95, 0x9C, 0x87, 0x8E, 0xB1, 0xB8, 0xA3, 0xAA,
	0xEC, 0xE5, 0xFE, 0xF7, 0xC8, 0xC1, 0xDA, 0xD3, 0xA4, 0xAD, 0xB6, 0xBF, 0x80, 0x89, 0x92, 0x9B,
	0x7C, 0x75, 0x6E, 0x67, 0x58, 0x51, 0x4A, 0x43, 0x34, 0x3D, 0x26, 0x2F, 0x10, 0x19, 0x02, 0x0B,
	0xD7, 0xDE, 0xC5, 0xCC, 0xF3, 0xFA, 0xE1, 0xE8, 0x9F, 0x96, 0x8D, 0x84, 0xBB, 0xB2, 0xA9, 0xA0,
	0x47, 0x4E, 0x55, 0x5C, 0x63, 0x6A, 0x71, 0x78, 0x0F, 0x06, 0x1D, 0x14, 0x2B, 0x22, 0x39, 0x30,
	0x9A, 0x93, 0x88, 0x81, 0xBE, 0xB7, 0xAC, 0xA5, 0xD2, 0xDB, 0xC0, 0xC9, 0xF6, 0xFF, 0xE4, 0xED,
	0x0A, 0x03, 0x18, 0x11, 0x2E, 0x27, 0x3C, 0x35, 0x42, 0x4B, 0x50, 0x59, 0x66, 0x6F, 0x74, 0x7D,
	0xA1, 0xA8, 0xB3, 0xBA, 0x85, 0x8C, 0x97, 0x9E, 0xE9, 0xE0, 0xFB, 0xF2, 0xCD, 0xC4, 0xDF, 0xD6,
	0x31, 0x38, 0x23, 0x2A, 0x15, 0x1C, 0x07, 0x0E, 0x79, 0x70, 0x6B, 0x62, 0x5D, 0x54, 0x4F, 0x46,
}
var g_DeBox4 = [...]uint8{
	0x00, 0x0B, 0x16, 0x1D, 0x2C, 0x27, 0x3A, 0x31, 0x58, 0x53, 0x4E, 0x45, 0x74, 0x7F, 0x62, 0x69,
	0xB0, 0xBB, 0xA6, 0xAD, 0x9C, 0x97, 0x8A, 0x81, 0xE8, 0xE3, 0xFE, 0xF5, 0xC4, 0xCF, 0xD2, 0xD9,
	0x7B, 0x70, 0x6D, 0x66, 0x57, 0x5C, 0x41, 0x4A, 0x23, 0x28, 0x35, 0x3E, 0x0F, 0x04, 0x19, 0x12,
	0xCB, 0xC0, 0xDD, 0xD6, 0xE7, 0xEC, 0xF1, 0xFA, 0x93, 0x98, 0x85, 0x8E, 0xBF, 0xB4, 0xA9, 0xA2,
	0xF6, 0xFD, 0xE0, 0xEB, 0xDA, 0xD1, 0xCC, 0xC7, 0xAE, 0xA5, 0xB8, 0xB3, 0x82, 0x89, 0x94, 0x9F,
	0x46, 0x4D, 0x50, 0x5B, 0x6A, 0x61, 0x7C, 0x77, 0x1E, 0x15, 0x08, 0x03, 0x32, 0x39, 0x24, 0x2F,
	0x8D, 0x86, 0x9B, 0x90, 0xA1, 0xAA, 0xB7, 0xBC, 0xD5, 0xDE, 0xC3, 0xC8, 0xF9, 0xF2, 0xEF, 0xE4,
	0x3D, 0x36, 0x2B, 0x20, 0x11, 0x1A, 0x07, 0x0C, 0x65, 0x6E, 0x73, 0x78, 0x49, 0x42, 0x5F, 0x54,
	0xF7, 0xFC, 0xE1, 0xEA, 0xDB, 0xD0, 0xCD, 0xC6, 0xAF, 0xA4, 0xB9, 0xB2, 0x83, 0x88, 0x95, 0x9E,
	0x47, 0x4C, 0x51, 0x5A, 0x6B, 0x60, 0x7D, 0x76, 0x1F, 0x14, 0x09, 0x02, 0x33, 0x38, 0x25, 0x2E,
	0x8C, 0x87, 0x9A, 0x91, 0xA0, 0xAB, 0xB6, 0xBD, 0xD4, 0xDF, 0xC2, 0xC9, 0xF8, 0xF3, 0xEE, 0xE5,
	0x3C, 0x37, 0x2A, 0x21, 0x10, 0x1B, 0x06, 0x0D, 0x64, 0x6F, 0x72, 0x79, 0x48, 0x43, 0x5E, 0x55,
	0x01, 0x0A, 0x17, 0x1C, 0x2D, 0x26, 0x3B, 0x30, 0x59, 0x52, 0x4F, 0x44, 0x75, 0x7E, 0x63, 0x68,
	0xB1, 0xBA, 0xA7, 0xAC, 0x9D, 0x96, 0x8B, 0x80, 0xE9, 0xE2, 0xFF, 0xF4, 0xC5, 0xCE, 0xD3, 0xD8,
	0x7A, 0x71, 0x6C, 0x67, 0x56, 0x5D, 0x40, 0x4B, 0x22, 0x29, 0x34, 0x3F, 0x0E, 0x05, 0x18, 0x13,
	0xCA, 0xC1, 0xDC, 0xD7, 0xE6, 0xED, 0xF0, 0xFB, 0x92, 0x99, 0x84, 0x8F, 0xBE, 0xB5, 0xA8, 0xA3,
}

func arrcopy(dst []uint8, src []uint8, dst_idx int, src_idx int, len int) {
	i := 0
	for i < len {
		dst[dst_idx] = src[src_idx]
		dst_idx++
		src_idx++
		i++
	}
}

type OblivionisAES struct {
	key []uint8
}

func ROR4(value uint32, count int) uint32 {
	count %= 32
	return (value >> count) | (value<<(32-count))&((1<<32)-1)
}

func MyXor(key []uint8, keyOffset int, data []uint8, dataOffset int) {
	i := 0
	for i < 16 {
		data[i+dataOffset] ^= key[i+keyOffset]
		i++
	}
}

func replaceBoxData(data []uint8, dataOffset int) {
	i := 0
	for i < 16 {
		data[i+dataOffset] = g_Box1[data[i+dataOffset]]
		i++
	}
}
func DeReplaceBoxData(data []uint8, dataOffset int) {
	i := 0
	for i < 16 {
		data[i+dataOffset] = g_DereplaceFourKeyBox[data[i+dataOffset]]
		i++
	}
}

func ByteOutOfOrder(data []uint8, dataOffset int) {
	ByteOut := make([]uint8, 16)
	ByteOut[0] = data[0+dataOffset]
	ByteOut[1] = data[5+dataOffset]
	ByteOut[2] = data[10+dataOffset]
	ByteOut[3] = data[15+dataOffset]
	ByteOut[4] = data[4+dataOffset]
	ByteOut[5] = data[9+dataOffset]
	ByteOut[6] = data[14+dataOffset]
	ByteOut[7] = data[3+dataOffset]
	ByteOut[8] = data[8+dataOffset]
	ByteOut[9] = data[13+dataOffset]
	ByteOut[10] = data[2+dataOffset]
	ByteOut[11] = data[7+dataOffset]
	ByteOut[12] = data[12+dataOffset]
	ByteOut[13] = data[1+dataOffset]
	ByteOut[14] = data[6+dataOffset]
	ByteOut[15] = data[11+dataOffset]
	arrcopy(data, ByteOut, dataOffset, 0, 16)
}
func DeByteOutOfOrder(data []uint8, dataOffset int) {
	ByteOut := make([]uint8, 16)
	ByteOut[0] = data[0+dataOffset]
	ByteOut[1] = data[13+dataOffset]
	ByteOut[2] = data[10+dataOffset]
	ByteOut[3] = data[7+dataOffset]
	ByteOut[4] = data[4+dataOffset]
	ByteOut[5] = data[1+dataOffset]
	ByteOut[6] = data[14+dataOffset]
	ByteOut[7] = data[11+dataOffset]
	ByteOut[8] = data[8+dataOffset]
	ByteOut[9] = data[5+dataOffset]
	ByteOut[10] = data[2+dataOffset]
	ByteOut[11] = data[15+dataOffset]
	ByteOut[12] = data[12+dataOffset]
	ByteOut[13] = data[9+dataOffset]
	ByteOut[14] = data[6+dataOffset]
	ByteOut[15] = data[3+dataOffset]
	arrcopy(data, ByteOut, dataOffset, 0, 16)
}

func BoxXorData(data []uint8, dataOffset int) {
	tmpData := make([]uint8, 16)

	// 第 1 个字节
	tmp := data[dataOffset+2] ^ data[dataOffset+3]
	tmp ^= g_Box3[data[dataOffset]]
	tmp ^= g_Box4[data[dataOffset+1]]
	tmpData[0] = tmp
	// 第 2 个字节
	tmp = data[dataOffset] ^ data[dataOffset+3]
	tmp ^= g_Box3[data[dataOffset+1]]
	tmp ^= g_Box4[data[dataOffset+2]]
	tmpData[1] = tmp
	// 第 3 个字节
	tmp = data[dataOffset] ^ data[dataOffset+1]
	tmp ^= g_Box3[data[dataOffset+2]]
	tmp ^= g_Box4[data[dataOffset+3]]
	tmpData[2] = tmp
	// 第 4 个字节
	tmp = data[dataOffset+1] ^ data[dataOffset+2]
	tmp ^= g_Box4[data[dataOffset+0]]
	tmp ^= g_Box3[data[dataOffset+3]]
	tmpData[3] = tmp

	// 第 5 个字节
	tmp = data[dataOffset+6] ^ data[dataOffset+7]
	tmp ^= g_Box3[data[dataOffset+4]]
	tmp ^= g_Box4[data[dataOffset+5]]
	tmpData[4] = tmp
	// 第 6 个字节
	tmp = data[dataOffset+4] ^ data[dataOffset+7]
	tmp ^= g_Box3[data[dataOffset+5]]
	tmp ^= g_Box4[data[dataOffset+6]]
	tmpData[5] = tmp
	// 第 7 个字节
	tmp = data[dataOffset+4] ^ data[dataOffset+5]
	tmp ^= g_Box3[data[dataOffset+6]]
	tmp ^= g_Box4[data[dataOffset+7]]
	tmpData[6] = tmp
	// 第 8 个字节
	tmp = data[dataOffset+5] ^ data[dataOffset+6]
	tmp ^= g_Box4[data[dataOffset+4]]
	tmp ^= g_Box3[data[dataOffset+7]]
	tmpData[7] = tmp

	// 第 9 个字节
	tmp = data[dataOffset+10] ^ data[dataOffset+11]
	tmp ^= g_Box3[data[dataOffset+8]]
	tmp ^= g_Box4[data[dataOffset+9]]
	tmpData[8] = tmp
	// 第 10 个字节
	tmp = data[dataOffset+8] ^ data[dataOffset+11]
	tmp ^= g_Box3[data[dataOffset+9]]
	tmp ^= g_Box4[data[dataOffset+10]]
	tmpData[9] = tmp
	// 第 11 个字节
	tmp = data[dataOffset+8] ^ data[dataOffset+9]
	tmp ^= g_Box3[data[dataOffset+10]]
	tmp ^= g_Box4[data[dataOffset+11]]
	tmpData[10] = tmp
	// 第 12 个字节
	tmp = data[dataOffset+9] ^ data[dataOffset+10]
	tmp ^= g_Box4[data[dataOffset+8]]
	tmp ^= g_Box3[data[dataOffset+11]]
	tmpData[11] = tmp

	// 第 13 个字节
	tmp = data[dataOffset+14] ^ data[dataOffset+15]
	tmp ^= g_Box3[data[dataOffset+12]]
	tmp ^= g_Box4[data[dataOffset+13]]
	tmpData[12] = tmp
	// 第 14 个字节
	tmp = data[dataOffset+12] ^ data[dataOffset+15]
	tmp ^= g_Box3[data[dataOffset+13]]
	tmp ^= g_Box4[data[dataOffset+14]]
	tmpData[13] = tmp
	// 第 15 个字节
	tmp = data[dataOffset+12] ^ data[dataOffset+13]
	tmp ^= g_Box3[data[dataOffset+14]]
	tmp ^= g_Box4[data[dataOffset+15]]
	tmpData[14] = tmp
	// 第 16 个字节
	tmp = data[dataOffset+13] ^ data[dataOffset+14]
	tmp ^= g_Box4[data[dataOffset+12]]
	tmp ^= g_Box3[data[dataOffset+15]]
	tmpData[15] = tmp

	arrcopy(data, tmpData, dataOffset, 0, 16)
}
func DeBoxXorData(data []uint8, dataOffset int) {
	tmpData := make([]uint8, 16)

	// 第 1 个字节
	tmp := g_DeBox1[data[dataOffset]]
	tmp ^= g_DeBox4[data[dataOffset+1]]
	tmp ^= g_DeBox2[data[dataOffset+2]]
	tmp ^= g_DeBox3[data[dataOffset+3]]
	tmpData[0] = tmp
	// 第 2 个字节
	tmp = g_DeBox3[data[dataOffset]]
	tmp ^= g_DeBox1[data[dataOffset+1]]
	tmp ^= g_DeBox4[data[dataOffset+2]]
	tmp ^= g_DeBox2[data[dataOffset+3]]
	tmpData[1] = tmp
	// 第 3 个字节
	tmp = g_DeBox2[data[dataOffset]]
	tmp ^= g_DeBox3[data[dataOffset+1]]
	tmp ^= g_DeBox1[data[dataOffset+2]]
	tmp ^= g_DeBox4[data[dataOffset+3]]
	tmpData[2] = tmp
	// 第 4 个字节
	tmp = g_DeBox4[data[dataOffset]]
	tmp ^= g_DeBox2[data[dataOffset+1]]
	tmp ^= g_DeBox3[data[dataOffset+2]]
	tmp ^= g_DeBox1[data[dataOffset+3]]
	tmpData[3] = tmp

	// 第 5 个字节
	tmp = g_DeBox1[data[dataOffset+4]]
	tmp ^= g_DeBox4[data[dataOffset+5]]
	tmp ^= g_DeBox2[data[dataOffset+6]]
	tmp ^= g_DeBox3[data[dataOffset+7]]
	tmpData[4] = tmp
	// 第 6 个字节
	tmp = g_DeBox3[data[dataOffset+4]]
	tmp ^= g_DeBox1[data[dataOffset+5]]
	tmp ^= g_DeBox4[data[dataOffset+6]]
	tmp ^= g_DeBox2[data[dataOffset+7]]
	tmpData[5] = tmp
	// 第 7 个字节
	tmp = g_DeBox2[data[dataOffset+4]]
	tmp ^= g_DeBox3[data[dataOffset+5]]
	tmp ^= g_DeBox1[data[dataOffset+6]]
	tmp ^= g_DeBox4[data[dataOffset+7]]
	tmpData[6] = tmp
	// 第 8 个字节
	tmp = g_DeBox4[data[dataOffset+4]]
	tmp ^= g_DeBox2[data[dataOffset+5]]
	tmp ^= g_DeBox3[data[dataOffset+6]]
	tmp ^= g_DeBox1[data[dataOffset+7]]
	tmpData[7] = tmp

	// 第 9 个字节
	tmp = g_DeBox1[data[dataOffset+8]]
	tmp ^= g_DeBox4[data[dataOffset+9]]
	tmp ^= g_DeBox2[data[dataOffset+10]]
	tmp ^= g_DeBox3[data[dataOffset+11]]
	tmpData[8] = tmp
	// 第 10 个字节
	tmp = g_DeBox3[data[dataOffset+8]]
	tmp ^= g_DeBox1[data[dataOffset+9]]
	tmp ^= g_DeBox4[data[dataOffset+10]]
	tmp ^= g_DeBox2[data[dataOffset+11]]
	tmpData[9] = tmp
	// 第 11 个字节
	tmp = g_DeBox2[data[dataOffset+8]]
	tmp ^= g_DeBox3[data[dataOffset+9]]
	tmp ^= g_DeBox1[data[dataOffset+10]]
	tmp ^= g_DeBox4[data[dataOffset+11]]
	tmpData[10] = tmp
	// 第 12 个字节
	tmp = g_DeBox4[data[dataOffset+8]]
	tmp ^= g_DeBox2[data[dataOffset+9]]
	tmp ^= g_DeBox3[data[dataOffset+10]]
	tmp ^= g_DeBox1[data[dataOffset+11]]
	tmpData[11] = tmp

	// 第 13 个字节
	tmp = g_DeBox1[data[dataOffset+12]]
	tmp ^= g_DeBox4[data[dataOffset+13]]
	tmp ^= g_DeBox2[data[dataOffset+14]]
	tmp ^= g_DeBox3[data[dataOffset+15]]
	tmpData[12] = tmp
	// 第 14 个字节
	tmp = g_DeBox3[data[dataOffset+12]]
	tmp ^= g_DeBox1[data[dataOffset+13]]
	tmp ^= g_DeBox4[data[dataOffset+14]]
	tmp ^= g_DeBox2[data[dataOffset+15]]
	tmpData[13] = tmp
	// 第 15 个字节
	tmp = g_DeBox2[data[dataOffset+12]]
	tmp ^= g_DeBox3[data[dataOffset+13]]
	tmp ^= g_DeBox1[data[dataOffset+14]]
	tmp ^= g_DeBox4[data[dataOffset+15]]
	tmpData[14] = tmp
	// 第 16 个字节
	tmp = g_DeBox4[data[dataOffset+12]]
	tmp ^= g_DeBox2[data[dataOffset+13]]
	tmp ^= g_DeBox3[data[dataOffset+14]]
	tmp ^= g_DeBox1[data[dataOffset+15]]
	tmpData[15] = tmp

	arrcopy(data, tmpData, dataOffset, 0, 16)
}

// 以下的函数是有用的。
/**
 * 获取一个 AES 加密/解密器
 * 参数是一个 16 字节的 uint8 数组, 是 AES 加/解的密钥
 */
func getAES(key []uint8) OblivionisAES {
	var i uint32

	g_key := make([]uint8, 176)
	tmp := make([]uint8, 4)
	arrcopy(g_key, key, 0, 0, 16)
	keyLen := 16
	keyIdx := 0
	replaceIdx := 1
	// 取 key 末尾 4 字节
	arrcopy(tmp, key, 0, 12, 4)
	for keyLen < 176 {
		if keyLen%16 == 0 { // 处理新一组的 16 个字节
			// 此时 i 用作一个临时暂存 tmp 数组的变量, 用来对 tmp 进行整体循环右移
			i = binary.LittleEndian.Uint32(tmp[0:4])
			i = ROR4(i, 8)
			binary.LittleEndian.PutUint32(tmp, i)
			tmp[0] = g_Box1[tmp[0]] ^ g_Box2[replaceIdx]
			replaceIdx++
			i = 1
			for i < 4 {
				tmp[i] = g_Box1[tmp[i]]
				i++
			}
		}
		i = 0 // 此时 i 用作数组下标
		for i < 4 {
			g_key[keyLen] = tmp[i] ^ g_key[keyIdx]
			tmp[i] = g_key[keyLen]
			keyLen++
			keyIdx++
			i++
		}
	}

	return OblivionisAES{
		key: g_key,
	}
}

/**
 * 利用 AES 加密器进行加密
 * 参数是需要进行加密的数据, 使用 uint8 数组的形式给出
 * 返回值就是加密后的数据
 */
func (aes OblivionisAES) EncryptData(pdData []uint8) []uint8 {
	dataIdx := 0
	originLen := len(pdData)
	dataLen := 16 * (len(pdData)/16 + 1)
	data := make([]uint8, dataLen)
	copy(data, pdData)
	padding := (uint8)(dataLen - originLen)
	j := 0
	for j < int(padding) {
		data[j+originLen] = padding
		j++
	}

	j = 0
	for j < dataLen {
		xorKeyIdx := 16
		MyXor(aes.key, 0, data, dataIdx)
		i := 0
		for i < 9 {
			replaceBoxData(data, dataIdx)
			ByteOutOfOrder(data, dataIdx)
			BoxXorData(data, dataIdx)
			MyXor(aes.key, xorKeyIdx, data, dataIdx)
			xorKeyIdx += 16
			i++
		}
		replaceBoxData(data, dataIdx)
		ByteOutOfOrder(data, dataIdx)
		MyXor(aes.key, xorKeyIdx, data, dataIdx)
		dataIdx += 16
		j += 16
	}
	return data
}

/**
 * 利用 AES 加密器进行解密
 * 参数是需要进行解密的数据, 使用 uint8 数组的形式给出
 * 返回值就是解密后的数据
 */
func (aes OblivionisAES) DecryptData(pdData []uint8) []uint8 {
	data := make([]uint8, len(pdData))
	copy(data, pdData)
	dataIdx := 0
	j := 0
	for j < len(data) {
		keyIdx := 0xA0
		MyXor(aes.key, keyIdx, data, dataIdx)
		DeByteOutOfOrder(data, dataIdx)
		DeReplaceBoxData(data, dataIdx)
		i := 0
		for i < 9 {
			keyIdx -= 0x10
			MyXor(aes.key, keyIdx, data, dataIdx)
			DeBoxXorData(data, dataIdx)
			DeByteOutOfOrder(data, dataIdx)
			DeReplaceBoxData(data, dataIdx)
			i++
		}
		MyXor(aes.key, 0, data, dataIdx)
		dataIdx += 16
		j += 16
	}

	trueLen := len(data) - int(data[len(data)-1])
	trueData := make([]uint8, trueLen)
	copy(trueData, data)
	return trueData
}

/**
 * 调试用函数
 * 参数是一个 uint8 数组
 * 这个函数将会把这个数组以 hexDump 的形式工整地打印在控制台上
 */
func hexDump(data []uint8) {
	i := 0
	for i < len(data) {
		if i != 0 && i%16 == 0 {
			fmt.Println("")
		}
		fmt.Printf("%02x ", data[i])
		i++
	}
	fmt.Println("")
}

// 测试用例
func test() {
	key := "1234567887654321"
	g_key := make([]uint8, len(key))
	bytes := []byte(key)
	for i, b := range bytes {
		g_key[i] = b
	}

	toEncrypt := "abcdefghabcdefghabcdefghabcdefgh1"
	encryData := make([]uint8, len(toEncrypt))
	bytes = []byte(toEncrypt)
	for i, b := range bytes {
		encryData[i] = b
	}
	aes := getAES(g_key)
	fmt.Println("to toEncrypt = ")
	hexDump(bytes)
	a := aes.EncryptData(bytes)
	fmt.Println("Encrypted = ")
	hexDump(a)
	b := aes.DecryptData(a)
	fmt.Println("Decrypted = ")
	hexDump(b)
}
