#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys

head = '''// 本代码由脚本根据{proto_file}和tmpl_api.go文件自动生成
// 请不要直接修改本文件

'''

class Api(object):
    fields = ['Method', 'Url', 'Param', 'Resp']
    def __init__(self, name):
        self.name = name
        self.fields = {}

def handle(filename):
    apis = {}
    pkg = None
    with open(filename) as f:
        for line in f:
            if not line.startswith('const ') and not line.startswith('type '):
                continue
            e = line.strip().split()
            if len(e) < 4:
                continue
            k = e[1]
            for item in Api.fields:
                if k.endswith(item):
                    name = k[:-len(item)]
                    if name not in apis:
                        apis[name] = Api(name)
                    apis[name].fields[item] = k

    h = head.format(proto_file=filename)
    t = open('tmpl_api.go').read()
    i = t.find('=====header====')
    if i > 0:
        j = t[:i].rfind('\n')
        h += t[:j]
        j = t.find('\n', i)
        t = t[j+1:]
    sys.stdout.write(h)
    for name in apis:
        api = apis[name]
        if len(api.fields) != len(Api.fields):
            sys.stderr.write('warnning: found {name}, but miss some fields'.format(name=name))
            continue
        s = t.replace('TmplName', name)
        sys.stdout.write(s)

if __name__ == '__main__':
    if len(sys.argv) != 2:
        sys.stderr.write('Usage:%s <xxx.proto.go>\n' % sys.argv[0])
        sys.exit(1)
    handle(sys.argv[1])
