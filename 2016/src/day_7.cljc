(ns day-7 (:require util))

(def input (util/read-input))

(def regex-a #"^(?!.*\[[a-z]*([a-z])([a-z])\2\1[a-z]*\].*)(?:[a-z\[\]]*([a-z])(?!\3)([a-z])\4\3[a-z\[\]]*)$")
(println (count (filter #(re-matches regex-a %) input)))

(def regex-b #".*([a-z])(?!\1)([a-z])\1[a-z]*(?:(?:[\[\]][a-z]+){2})*[\[\]][a-z]*\2\1\2.*")
(println (count (filter #(re-matches regex-b %) input)))
