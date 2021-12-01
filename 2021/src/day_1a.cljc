(ns day-1a
  (:require [clojure.string :as str]))

(def input
  (map #(Integer/parseInt %) (str/split (slurp "../inputs/1/real.txt") #"\n")))

(def answer
  (count (filter identity (map < input (rest input)))))

(println answer)
