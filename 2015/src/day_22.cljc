(ns day-22)

(def spells [{:name "Magic Missile" :price 53 :instantly #(update % :boss-hit-points + -4)}
             {:name "Drain" :price 73 :instantly #(update (update % :boss-hit-points + -2) :player-hit-points + 2)}
             {:name "Shield" :price 113 :lasts-for 6 :effect #(update % :player-armor + 7)}
             {:name "Poison" :price 173 :lasts-for 6 :effect #(update % :boss-hit-points + -3)}
             {:name "Recharge" :price 229 :lasts-for 5 :effect #(update % :player-mana + 101)}])

(defn apply-effects
  [{:keys [active-effects] :as state}]
  (assoc
    (reduce #((:effect %2) %1) state active-effects)
    :active-effects
    (remove (comp zero? :lasts-for) (map #(update % :lasts-for dec) active-effects))))

(defn boss-turn
  [state]
  (let [{:keys [boss-damage player-armor] :as after-effects} (apply-effects state)]
    (-> after-effects
        (update :player-hit-points #(- % (max 1 (- boss-damage player-armor))))
        (assoc :player-armor 0))))

(defn cast-spell
  [state spell]
  (let [after-paying (update (update state :player-mana + (unchecked-negate (:price spell))) :spent-mana + (:price spell))]
    (if (contains? spell :instantly)
      ((:instantly spell) after-paying)
      (update after-paying :active-effects conj spell))))

(defn player-turn
  [state pre-turn-hp]
  (let [{:keys [active-effects] :as after-effects} (apply-effects (update state :player-hit-points - pre-turn-hp))]
    (map
     #(-> after-effects
          (cast-spell %)
          (assoc :player-armor 0))
     (remove
       (some-fn
         #((set (map :name active-effects)) (:name %))
         #(> (:price %) (:player-mana after-effects)))
       spells))))

(def initial-state {:boss-hit-points 55 :boss-damage 8 :player-hit-points 50 :player-armor 0 :player-mana 500 :spent-mana 0})

(defn find-cheapest-win
  [state pre-turn-hp max-search]
  (reduce (fn [max-search after-player-turn]
            (cond
              (>= (:spent-mana after-player-turn) max-search)
              max-search
              (<= (:boss-hit-points after-player-turn) 0)
              (:spent-mana after-player-turn)
              :else
              (let [after-boss-turn (boss-turn after-player-turn)]
                (if (<= (:player-hit-points after-boss-turn) pre-turn-hp)
                  max-search
                  (find-cheapest-win after-boss-turn pre-turn-hp max-search)))))
          max-search
          (player-turn state pre-turn-hp)))

(println (find-cheapest-win initial-state 0 Integer/MAX_VALUE))
(println (find-cheapest-win initial-state 1 Integer/MAX_VALUE))
