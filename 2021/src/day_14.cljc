(ns day-14 (:require util))

(def input
  (util/read-input))

(def pair-insertion-matrix
  (let [insertion-pairs (map #(clojure.string/split % #" -> ") (drop 2 input))]
    (reduce
      (fn
        [m [from to]]
        (update-in m (apply vector (seq from)) (fn [_] [(first to) (second from)])))
      {}
      insertion-pairs)))

(defn expand-pair
  [& args]
  (get-in pair-insertion-matrix args))

(defn expand-polymer
  [s]
  (cons (first s) (mapcat expand-pair s (rest s))))

(defn expanded-polymer*
  [s]
  (lazy-seq (cons s (expanded-polymer* (expand-polymer s)))))

(def expanded-polymer
  (-> input first seq expanded-polymer*))

(def after-step-10
  (frequencies (nth expanded-polymer 10)))

(def result-part-a
  (apply - (map (comp second #(apply % second after-step-10)) [max-key min-key])))

(println result-part-a)
