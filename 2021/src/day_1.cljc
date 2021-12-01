(ns day-1
  (:require [clojure.string :as str]))

(def input
  (map #(Integer/parseInt %) (str/split (slurp "../inputs/1/real.txt") #"\n")))

(defn sliding-sums
  [n xs]
  (map #(apply + %) (partition n 1 xs)))

(defn count-increasing
  [xs]
  (count (filter identity (map < xs (rest xs)))))

(println (count-increasing input))
(println (count-increasing (sliding-sums 3 input)))
