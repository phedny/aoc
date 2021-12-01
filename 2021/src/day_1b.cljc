(ns day-1b
  (:require [clojure.string :as str]))

(def input
  (map #(Integer/parseInt %) (str/split (slurp "../inputs/1/real.txt") #"\n")))

(def windows
  (map + input (rest input) (rest (rest input))))

(def answer
  (count (filter identity (map < windows (rest windows)))))

(println answer)
