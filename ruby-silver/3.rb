a = [0, 1, 2, 3, 4, 5]
ra = a.delete_if{|x| x % 2 == 0}
p 'a:',  a #->[1, 3, 5]
p 'ra:', ra #->[1, 3, 5]


b = [0, 1, 2, 3, 4, 5]
br = b.reject!
br.each{|i| i % 2 == 0}
p 'b:', b #->[1, 3, 5]
p 'br:', br #->#<Enumerator: [1, 3, 5]:reject!>


c = [10]
cr = c.reject!
cr.each{|j| j % 2 == 0}
p 'c:', c #->[]
p 'cr:', cr #->#<Enumerator: []:reject!>


d = [10]
rd = d.delete_if{|x| x % 2 == 0}
p 'd:', d #->[]
p 'rd:', rd #->[]


e = [9]
er = e.delete_if{|x| x % 2 == 0}
p 'e:',  e #->[9]
p 'er:', er #->[9]

f = [9]
fr = c.reject!
fr.each{|j| j % 2 == 0}
p 'f:', f #->[9]
p 'fr:', fr #->#<Enumerator: []:reject!>
