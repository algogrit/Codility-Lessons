#!/usr/bin/env ruby

def get_friends(network, list_of_friends, person)
  first_degree_friends = network[person]
  if first_degree_friends.nil? || first_degree_friends.empty?
    return list_of_friends
  end
  first_degree_friends.reduce(list_of_friends) do |list, friend|
    if list.include? friend
      list
    else
      list << friend
      get_friends(network, list, friend)
    end
  end
end

def number_of_friends(relations: , people:)
  network = {}

  relations.each do |relation|
    friend, other_friend = relation.split(":")

    network[friend] = (network[friend] || []) << other_friend
    network[other_friend] = (network[other_friend] || []) << friend
  end

  people.map do |person|
    inclusive_friends = get_friends(network, [person], person)
    friends = inclusive_friends - [person]

    {person => [friends.count, friends.join(", ")]}
  end
end

# Anne:2, Vinny: 1
pp number_of_friends(relations: ["Anne:Barbara", "Barbara:Ivan", "Vinny:Vlad"], people: ["Anne", "Vinny"])
# Octavia: 4, Diego: 4
pp number_of_friends(relations: ["Octavia:Zoltan", "Zoltan:Marko", "Marko:Diego", "Diego:Andres"], people: ["Octavia", "Diego"])

# Tomislav: 3, Martin: 4
pp number_of_friends(relations: ["Vanja:Sergio", "Sergio:Pedro", "Pedro:Martin", "Martin:Pablo", "Pablo:Sergio", "Jelena:Ivan", "Jelena:Alan", "Alan:Tomislav"], people: ["Tomislav", "Martin"])

# Zoltan:1, Ivan: 1
pp number_of_friends(relations: ["Alvaro:Alex", "Roman:Nikola", "Octavia:Barbara", "Joao:Marko", "Luis:Vanja", "Gabriel:Gustavo", "Alan:Pablo", "Ivan:Andres", "Artem:Anne", "Martin:Alessandro", "Sebastian:Vinny", "Eduardo:Francis", "Zoltan:Vlad"], people: ["Zoltan", "Ivan"])

# Francis: 5, David: 5
pp number_of_friends(relations: ["David:Francis", "Francis:Alessandro", "Alessandro:Ivan", "Tomislav:Ivan", "Anne:Tomislav", "Anne:David", "Francis:Tomislav"], people: ["Francis", "David"])

# Barbara: 6, Joao: 7
pp number_of_friends(relations: ["Alessandro:Anna", "Anna:Anne", "Anne:Barbara", "Barbara:David", "David:Francis", "Francis:Eduardo", "Eduardo:Anna", "Eduardo:Alessandro", "Luis:Marko", "Joao:Vlad", "Vlad:Luka", "Luka:Nikola", "Nikola:Roman", "Vlad:Roman", "Vlad:Vinny", "Vinny:Roman", "Vlad:Andres", "Vinny:Ivan"], people: ["Barbara", "Joao"])
