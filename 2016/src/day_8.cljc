(ns day-8 (:require util [clojure.string :as str]))

(def input (util/read-input #"\n" #(str/split % #"(?: (?:by )?)|=|(?<=\d)x(?=\d)")))

(defn apply-line
  [on [instr a0 a1 a2 a3]]
  (case instr
    "rect" (reduce conj on (for [x (range (parse-long a0))
                                 y (range (parse-long a1))]
                             [x y]))
    "rotate" (case a0
               "column" (set (map (fn [[x y :as c]]
                                    (if (= x (parse-long a2))
                                      [x (mod (+ y (parse-long a3)) 6)]
                                      c))
                                  on))
               "row" (set (map (fn [[x y :as c]]
                                 (if (= y (parse-long a2))
                                   [(mod (+ x (parse-long a3)) 50) y]
                                   c))
                               on)))))

(def screen (reduce apply-line #{} input))

(println (count screen))

(doseq [line (partition 50 (for [y (range 6)
                                 x (range 50)]
                             (if (contains? screen [x y]) \# \.)))]
  (println (apply str line)))
