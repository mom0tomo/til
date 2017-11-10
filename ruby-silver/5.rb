#hoge = a.to_h
#puts hoge.class

fuga = [["Alice", 50], ["Bob", 40], ["Charlie", 70]].to_h
piyo = [[:Alice, 50], [:Bob, 40], [:Charlie, 70]].to_h

p fuga
p piyo

p piyo[:Alice], piyo[':Alice'], fuga['Alice']
puts fuga.class

piyoa = piyo.to_a
p piyoa
p piyoa.class

p piyo.keys
p piyo.values
