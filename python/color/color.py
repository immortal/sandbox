codeCodes = {
    'black': u'0;30', 'bright gray': u'0;37',
    'blue': u'0;34', 'white': u'1;37',
    'green': u'0;32', 'bright blue': u'1;34',
    'cyan': u'0;36', 'bright green': u'1;32',
    'red': u'0;31', 'bright cyan': u'1;36',
    'purple': u'0;35', 'bright red': u'1;31',
    'yellow': u'0;33', 'bright purple': u'1;35',
    'dark gray': u'1;30', 'bright yellow': u'1;33',
    'magenta': u'0;35', 'bright magenta': u'1;35',
    'normal': u'0',
}


def stringc(text, color):
    return "\n".join([u"\033[%sm%s\033[0m" % (codeCodes[color], t) for t in text.split('\n')])

for k in codeCodes:
    print stringc("The quick brown fox jumps over the lazy dog", k), " -- ", k
