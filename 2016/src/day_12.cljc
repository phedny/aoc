(ns day-12 (:require [clojure.string :as str] util))

(def input (vec (util/read-input (fn [line]
                                   (let [[instr x y] (str/split line #" ")
                                         parse-parameter #(cond
                                                            (nil? %) nil
                                                            (re-matches #"-?\d+" %) (parse-long %)
                                                            :else (keyword %))]
                                     {:instr (keyword instr) :x (parse-parameter x) :y (parse-parameter y)})))))

(defn step
  [program state]
  (let [{:keys [instr x y]} (get program (:ip state))]
    (case instr
      :cpy (-> state (assoc y (if (number? x) x (x state))) (update :ip inc))
      :inc (-> state (update x inc) (update :ip inc))
      :dec (-> state (update x dec) (update :ip inc))
      :jnz (if (zero? (if (number? x) x (x state))) (update state :ip inc) (update state :ip + y)))))

(defn execute
  [program state]
  (if (< (:ip state) (count program))
    (recur program (step program state))
    state))

(println (:a (execute input {:ip 0 :a 0 :b 0 :c 0 :d 0})))
(println (:a (execute input {:ip 0 :a 0 :b 0 :c 1 :d 0})))
