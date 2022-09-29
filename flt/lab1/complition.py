import sys, re
from copy import copy, deepcopy

from models.trs import Trs
from models.term import Term
from models.graph import Graph


class Catapult(Exception):
    pass


class Algorithm():
    def __init__(self, order, rules):
        if order == 'anti-lexicographic':
            self.order = -1
        else:
            self.order = 1
        self.rules = set(rules)

    def start(self):
        try:
            self.knuth_bendix(Graph([]), self.rules)
        except Catapult:
            print("Finished")
        else:
            print("Unable to find order")

    def knuth_bendix(self, lex: Graph, rules: list[Trs]):
        lex = copy(lex)
        rules = copy(rules)
        for rule in rules:
            if rule.right_term == rule.left_term:
                return
        
        if lex.is_cyclic():
            # print("cycle found", lex)
            return
        
        if len(rules) == 0:
            print("Answer:")
            for v in lex.graph:
                for u in lex.graph[v]:
                    print(f"{v} >lg {u}")
            raise Catapult

        rule = rules.pop()

        if str(rule.left_term) in str(rule.right_term):
            return
        
        if not(self.kb1(rule, lex, rules)):
            self.kb3(rule, lex, rules) 
            self.kb4(rule, lex, rules)
            self.kb2(rule, lex, rules)

    def kb1(self, rule, lex, rules):
        if str(rule.right_term) in str(rule.left_term):
            self.knuth_bendix(lex, rules)
            return True
        return False

    def kb2(self, rule, lex, rules):
        if rule.left_term.ttype == 'constructor' and rule.right_term.ttype == 'constructor':
            for term in rule.left_term.targs:
                rules.add(Trs(term, rule.right_term))
                self.knuth_bendix(lex, rules)

    def kb3(self, rule, lex, rules):
        if rule.left_term.ttype == 'constructor' and rule.right_term.ttype == 'constructor':
            if rule.left_term.tname == rule.right_term.tname:
                return
            # print(lex)
            lex.add(rule.left_term.tname, rule.right_term.tname)
            # print(lex)
            for term in rule.right_term.targs:
                rules.add(Trs(rule.left_term, term))
            self.knuth_bendix(lex, rules)
    
    def kb4(self, rule, lex, rules):
        if rule.left_term.ttype == 'constructor' and rule.right_term.ttype == 'constructor':
            if rule.left_term.tname != rule.right_term.tname:
                return
            
            for term in rule.right_term.targs:
                rules.add(Trs(rule.left_term, term))
            
            for left_term, right_term in [*zip(rule.left_term.targs, rule.right_term.targs)][::self.order]:
                if left_term != right_term:
                    rules.add(Trs(left_term, right_term))
                    self.knuth_bendix(lex, rules)