from curses import raw
import re, sys
from models.term import Term
from models.trs import Trs


class Parser():
    def __init__(self, filepath):
        self.filepath = filepath
        
        with open(filepath, 'r') as f:
            content  = f.read().strip()
        
        pattern = re.compile(r"""
                                \s*(?P<order>.*?)
                                \s*constructors\s*=\s*(?P<constructors>(?:(\w|\s|,|\(|\)))*)
                                \s*variables\s*=\s*(?P<variables>(?:(\s|\w|,)*))
                                \s*(?P<trs>.*)
                                """, re.VERBOSE)
        content = re.sub('\s+', ' ', content)
        match = re.match(pattern, content)
        self.order = match.group("order")
        
        constructors_raw = match.group("constructors")
        self.constructors = self.parse_constructors(constructors_raw)
        
        variables_raw = match.group("variables")
        self.variables = self.parse_variables(variables_raw)

        trs_raw = match.group("trs")
        self.trss = self.parse_trs(trs_raw)

        #return order, constructors, variables, terms

    def parse_constructors(self, constructors_raw: str):
        constructors_raw = constructors_raw.split(",")
        constructors = {}
        for constructor in constructors_raw:
            tmp = constructor.strip()
            constructors[tmp[0]] = int(tmp[2])
        return constructors

    def parse_variables(self, variables_raw: str):
        variables_raw = variables_raw.split(",")
        variables = []
        for variable in variables_raw:
            variables.append(variable.strip())
        return variables

    def parse_trs(self, trs_raw):
        trss = []
        trs_raw = trs_raw.strip()
        while trs_raw:
            i = self.get_zero_level_skobka_index(trs_raw)
            trss.append(trs_raw[1:i])
            trs_raw = trs_raw[i+1:]
            trs_raw = trs_raw.strip()
        trss = [*map(self.parse_trs2, trss)]
        # for i in trss:
        #     print(i)
        return trss

    def parse_trs2(self, rule: str):
        # print(f"parse_trs2: {rule}")
        rule = rule.split("=")
        l, r = rule[0], rule[1]
        l = l.strip()
        r = r.strip()
        l, r = self.parse_term(l), self.parse_term(r)
        return Trs(l, r)

    def parse_term(self, term_raw: str):
        # print(f"parse_term: {term_raw}")
        term_raw = term_raw.strip()
        if term_raw[0] in self.variables:
            return Term('variable', None, term_raw[0])
        return Term('constructor', self.get_raw_subterms(term_raw), term_raw[0])

    def get_raw_subterms(self, raw_term):
        # print(f"get_raw_subterms: {raw_term}")
        # print(raw_term)
        raw_term = raw_term.strip()
        raw_term = raw_term[2:]
        # print(raw_term)
        subterms = []
        while raw_term:
            if raw_term[0] in self.variables:
                subterms.append(Term('variable', None, raw_term[0]))
                raw_term = raw_term[1:]
                for x in raw_term:
                    if not str.isalpha(x):
                        raw_term = raw_term[1:]
                    else: 
                        break
            elif raw_term[0] in self.constructors.keys():
                i = self.get_zero_level_skobka_index(raw_term)
                subterms.append(self.parse_term(raw_term[:i + 1]))
                raw_term = raw_term[i + 1:]
                for x in raw_term:
                    if not str.isalpha(x):
                        raw_term = raw_term[1:]
                    else: 
                        break
            else:
                print(raw_term)
                raise Exception("parsing error")
        return subterms

    
    def get_zero_level_skobka_index(self, stroka: str):
        counter = 0
        for i in range(len(stroka)):
            if stroka[i] == '(':
                counter += 1
            elif stroka[i] == ')':
                counter -= 1
            if counter == 0 and i != 0:
                return i
        return -1

    def get_parsed_data(self):
        return self.order, self.variables, self.constructors, self.trss
