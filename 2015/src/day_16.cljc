(ns day-16 (:require util))

(def input
  (util/read-input (juxt
                     #(second (first (re-seq #"Sue (\d+):" %)))
                     #(into {} (map (fn [[_ i c]] [(keyword i) (Long/parseLong c)]) (re-seq #"(\w+): (\d+)" %))))))

(def results
  {:children 3 :cats 7 :samoyeds 2 :pomeranians 3 :akitas 0 :vizslas 0 :goldfish 5 :trees 3 :cars 2 :perfumes 1})

(defn matches?
  [match-fns sue]
  (every? #((get match-fns (first %) =) (second %) (get results (first %))) sue))

(defn find-match
  [match-fns]
  (some #(if (matches? match-fns (second %)) (first %)) input))

(println (find-match {}))
(println (find-match {:cars > :trees > :pomeranians < :goldfish <}))
