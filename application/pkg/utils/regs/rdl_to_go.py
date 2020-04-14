#!/usr/bin/env python3

import sys

from systemrdl import RDLCompiler, RDLListener, RDLWalker, RDLCompileError
from systemrdl.node import FieldNode

import argparse

################################################################################
class RdlToGoCodeListener(RDLListener):
    rnames = ""
    rnames_counter = 0
    r32 = ""
    name = ""

    def enter_Addrmap(self, node):
        #self.name = node.type_name
        self.r32 += "var Registers = [...]register32 {\n"

    def exit_Addrmap(self, node):
        self.r32 += "}\n"

    def enter_Reg(self, node):
        self.r32 += "register32 {\n"
        self.r32 += "addr: " + hex(node.absolute_address) + ",\n"
        self.r32 += 'name: "' + node.type_name.upper() + '",\n'
        self.r32 += 'desc: "' + node.get_property("name") + '",\n'
        self.r32 += "fields: []field {\n"

        self.rnames += node.type_name.upper() + '= ' + str(self.rnames_counter) + '\n'
        self.rnames_counter = self.rnames_counter + 1

    def exit_Reg(self, node):
        self.r32 += "},\n"
        self.r32 += "},\n"

    def enter_Field(self, node):
        self.r32 += "field {\n"
        self.r32 += "bitStart: " + str(node.low) + ",\n"
        self.r32 += "bitEnd: " + str(node.high) + ",\n" 
        self.r32 += 'name: "' + node.type_name  + '",\n'
        self.r32 += 'desc: "' + node.get_property("name") + '",\n'
        if (node.get_property("encode")):
            self.r32 += "values: []value {\n"
            for item in node.get_property("encode"):
                self.r32 += "value {\n"
                name = str(item).split(".")
                self.r32 += 'name: "' + name[1] + '",\n'
                self.r32 += "value: " + str(item.value) + ",\n"
                self.r32 += "},\n"
            self.r32 += "},\n"

    def exit_Field(self, node):
        self.r32 += "},\n"

################################################################################
parser = argparse.ArgumentParser(description='dfdfdfd.')

parser.add_argument('rdl',     type=str,       help='SystemRDL source file')
parser.add_argument('chip',    type=str,       help='Chip model name')

parser.add_argument('--tags', type=str,   help='GO build tags',               required=False)

args = parser.parse_args()

#print(args)

rdlc = RDLCompiler()

try:
    rdlc.compile_file(args.rdl)

    root = rdlc.elaborate(args.chip)
except RDLCompileError:
    sys.exit(1)

#sys.exit(10)

walker = RDLWalker(unroll=True)
golistener = RdlToGoCodeListener()
walker.walk(root, golistener)

print("//THIS FILE WAS AUTO GENERATED. PLEASE DO NOT EDIT")

if args.tags != None:
    print("//+build " + str(args.tags))

print("package regs")

#print('func init() { addAddrMap("%s", %sRegisters[:]) }' % (golistener.name, golistener.name))

print('const (')
print(golistener.rnames)
print(')')

print(golistener.r32)
