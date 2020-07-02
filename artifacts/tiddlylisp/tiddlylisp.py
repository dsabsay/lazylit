# Commit: 2e5d71f9cfe572379f26713a2cd7ba77e9d361f2
# CommitDate: Mon May 25 2020
# SourceFile: tiddlylisp.py

"""
Most (or all) of this code is from Michael Nielsen's essay found here:
http://www.michaelnielsen.org/ddi/lisp-as-the-maxwells-equations-of-software/
which, in turn, was inspired by Peter Norvig's Python Lisp interpreter.
"""
import sys
import traceback
import readline


#### Data structures and types ####################

Symbol = str

class Env(dict):
    """
    An environment: Dictionary of {'var': val} pairs, with an outer
    Environment.
    """
    def __init__(self, params=(), args=(), outer=None):
        self.update(zip(params, args))
        self.outer = outer

    def find(self, var):
        """ Return innermost Env where var appears. """
        try:
            return self if var in self else self.outer.find(var)
        except AttributeError:
            raise Exception(f'Symbol "{var}" is unbound!')


def add_globals(env):
    """ Adds built-in procedures and variables to env. """
    import operator

    def tl_assert(x):
        if not x:
            raise AssertionError('Assertion failed!')

    env.update({
        '+': operator.add,
        '-': operator.sub,
        '*': operator.mul,
        '/': operator.truediv,
        '>': operator.gt,
        '<': operator.lt,
        '>=': operator.ge,
        '<=': operator.le,
        '=': operator.eq,
        'assert': tl_assert,
    })
    env.update({'True': True, 'False': False})
    return env


global_env = add_globals(Env())


def tl_eval(x, env=global_env):
    """ Evaluate an expression *x* in an environment *env*. """
    print(f'x: {x}')
    if isinstance(x, Symbol):  # variable reference
        return env.find(x)[x]
    elif not isinstance(x, list):  # constant literal
        return x
    elif x[0] == 'quote' or x[0] == 'q':
        _, exp = x
        return exp
    elif x[0] == 'atom?':
        _, exp = x
        return not isinstance(tl_eval(exp, env), list)
    elif x[0] == 'eq?':
        _, exp1, exp2 = x
        v1, v2 = tl_eval(exp1, env), tl_eval(exp2, env)
        if isinstance(v1, list):
            return len(v1) == 0 and v1 == v2
        else:
            return v1 == v2
    elif x[0] == 'car':
        _, exp = x
        # NOTE: This will evaluate all items of the list.
        return tl_eval(exp, env)[0]
    elif x[0] == 'cdr':
        _, exp = x
        return tl_eval(exp, env)[1:]
    elif x[0] == 'cons':
        _, exp1, exp2 = x
        return [tl_eval(exp1, env), *tl_eval(exp2, env)]
    elif x[0] == 'cond':
        for p, e in x[1:]:
            # This was causing problems because it made Python str type objects
            # behave like a Lisp boolean value:
            if tl_eval(p, env):
            # So we change to stricter comparison:
            # if tl_eval(p, env) is True:
            # But then one of Michael's tests failed, so I reverted it.
            # See NOTE1 in notex.txl.
                return tl_eval(e, env)
    elif x[0] == 'null?':
        _, exp = x
        return tl_eval(exp, env) == []
    elif x[0] == 'if':
        _, test, conseq, alt = x
        if tl_eval(test, env):
            return tl_eval(conseq, env)
        else:
            return tl_eval(alt, env)
    elif x[0] == 'define':
        _, var, exp = x
        env[var] = tl_eval(exp, env)
    elif x[0] == 'set!':
        _, var, exp = x
        env.find(var)[var] = tl_eval(exp, env)
    elif x[0] == 'lambda':
        _, params, exp = x
        return lambda *args: tl_eval(exp, Env(params, args, env))
    elif x[0] == 'begin':
        for exp in x[1:]:
            val = tl_eval(exp, env)
        return val
    else:  # procedure application
        exps = [tl_eval(exp, env) for exp in x]
        proc = exps.pop(0)
        return proc(*exps)


#### Parsing ############################################

def tl_parse(s):
    """ Parse a TL epxression from a string. """
    return read_from(tokenize(s))


def tokenize(s):
    """ Converts a string to list of tokens. """
    return s.replace('(', ' ( ').replace(')', ' ) ').split()


def read_from(tokens):
    """ Read an expression from a sequence of tokens. """
    if len(tokens) == 0:
        raise SyntaxError('Unexpected EOF while reading.')
    token = tokens.pop(0)
    if token == '(':
        L = []
        while tokens[0] != ')':
            L.append(read_from(tokens))
        tokens.pop(0)  # pop off ')'
        return L
    elif token == ')':
        raise SyntaxError('Unexpected )')
    else:
        return atom(token)


def atom(token):
    """ Convert a token to an atom. """
    try:
        return int(token)
    except ValueError:
        try:
            return float(token)
        except ValueError:
            return Symbol(token)


def tl_to_string(exp):
    """ Converts Python objects to a TL-readable string. """
    if not isinstance(exp, list):
        # return str(exp)
        return str(exp) + ' ' + str(type(exp))
    return '(' + ' '.join(map(tl_to_string, exp)) + ')'


#### Loading and running from file ####################

def running_paren_sums(program):
    """
    Map the lines in the list *program* to a list whose entries contain
    a running sum of the per-line difference between the number of '('
    and the number of ')'.
    """
    count_open_parens = lambda line: line.count('(') - line.count(')')
    paren_counts = map(count_open_parens, program)
    rps = []
    total = 0
    for paren_count in paren_counts:
        total += paren_count
        rps.append(total)
    return rps


def load(filename):
    """
    Load the tiddlylisp program in *filename*, execute it and start the REPl.
    """
    print(f'Loading and executing {filename}')
    with open(filename, 'r') as f:
        program = f.readlines()
    # Remove full-line comments
    program = [x for x in program if not x.lstrip().startswith(';')]
    rps = running_paren_sums(program)
    full_line = ''
    for paren_sum, program_line in zip(rps, program):
        if paren_sum < 0:
            raise SyntaxError(f'Too many ")" on line: {program_line}')
        program_line = program_line.strip()
        full_line += program_line + ' '
        if paren_sum == 0 and full_line.strip() != '':
            try:
                val = tl_eval(tl_parse(full_line))
                if val is not None:
                    print(tl_to_string(val))
            except:
                handle_error()
                print(f'The line in which the error occurred:\n{program_line}\n')
                break
            full_line = ''
    repl()


#### REPL ########################################

def repl(prompt='tiddlylisp> '):
    while True:
        try:
            val = tl_eval(tl_parse(input(prompt)))
            if val is not None:
                print(tl_to_string(val))
        except KeyboardInterrupt:
            print('\nExiting.\n')
            sys.exit()
        except:
            handle_error()


def handle_error():
    print('An error occurred. Python stacktrace:')
    traceback.print_exc()


if __name__ == '__main__':
    if len(sys.argv) == 2:
        load(sys.argv[1])
    else:
        repl()
