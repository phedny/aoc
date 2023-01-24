(ns day-23 (:require util))

(def input
  (mapv
    (fn [[_ instr p1w p1n p2n]]
      {:instr (keyword instr)
       :p1 (or (when (some? p1w) (keyword p1w)) (Integer/parseInt p1n))
       :p2 (when (some? p2n) (Integer/parseInt p2n))})
    (util/read-input #(first (re-seq #"(\w+) (?:(\w)|([+-]\d+))(?:, ([+-]\d+))?" %)))))

(defn step
  [program state]
  (when (> (count program) (:ip state))
    (let [{:keys [instr p1 p2]} (program (:ip state))]
      (case instr
        :hlf (-> state (update p1 / 2) (update :ip inc))
        :tpl (-> state (update p1 * 3) (update :ip inc))
        :inc (-> state (update p1 inc) (update :ip inc))
        :jmp (update state :ip + p1)
        :jie (if (even? (state p1)) (update state :ip + p2) (update state :ip inc))
        :jio (if (= (state p1) 1) (update state :ip + p2) (update state :ip inc))))))

(defn steps*
  [program state]
  (if (map? state)
    (lazy-seq (cons state (steps* program (step program state))))
    nil))

(println (:b (last (steps* input {:a 0 :b 0 :ip 0}))))
(println (:b (last (steps* input {:a 1 :b 0 :ip 0}))))
