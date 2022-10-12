(ns day-6 (:require util))

(def input (util/read-input))

(defn transpose
  [xs]
  (apply map vector xs))

(defn find-character
  [k s]
  (let [freqs (frequencies s)]
    (apply k freqs (keys freqs))))

(println (apply str (map (partial find-character max-key) (transpose input))))

(println (apply str (map (partial find-character min-key) (transpose input))))
