functor
import
  System
define
  proc {Mcm Chain}
    local
      N = {Length Chain}
      Aux = %{(i, i): (0,) + chain[i] for i in range(n)}
    in
      % Single matrix chain has zero cost
      for I in 1..N do % I: Length of subchain
        for J in 0..(N - I) do % I: Starting index of subchain
          local
            Best = 999999999 % Infinity
          in
            for K in J..(J + I) do % K: Splitting point of subchain
              local
                % Multiply subchains at splitting point
                %LCost, LName, LRow, LCol = Aux[J, K]
                %RCost, RName, RRow, RCol = Aux[K + 1, J + I]
                %Cost = LCost + RCost + LRow * LCol * RCol
                Var = "("#LName#RName#")"
              in
                % Pick the best one
                if Cost < Best then
                  Best = Cost
                  %Aux[J, J + I] = Cost, Var, LRow, RCol
                end
              end
            end
          end
        end
      end
    end
  end

  {Mcm [tree("A" 10 20)
        tree("B" 20 30)
        tree("C" 30 40)]}
end
