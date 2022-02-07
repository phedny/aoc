(ns day-19 (:require util [clojure.pprint :refer [pprint]]))

(def replacements
  (apply merge-with concat
         (util/read-input (comp (fn [[[_ from to]]] {(keyword from) [(map keyword (re-seq #"[A-Z][a-z]?" to))]}) (partial re-seq #"(\w+) => (\w+)")))))

(def medicine
  (map keyword (re-seq #"[A-Z][a-z]?" "CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl")))

(defn step
  [molecule]
  (mapcat
    (fn [i]
      (let [prefix (take i molecule)
            suffix (drop (inc i) molecule)]
        (map #(concat prefix % suffix) (replacements (nth molecule i)))))
    (range (count molecule))))

(def result-part-a
  (count (set (step medicine))))

(println result-part-a)

(defn count-of
  [c molecule]
  (count (filter #{c} molecule)))

(def result-part-b
  (- (count medicine) 1 (* 2 (count-of :Rn medicine)) (* 2 (count-of :Y medicine))))

(println result-part-b)

