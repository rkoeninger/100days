functor
import
  System
define
  fun {Range N M}
    if N == M
      then nil
      else N | {Range (N + 1) M}
    end
  end

  fun {Swap Xs N F}
    case Xs of
    X | Rest then
      if N == 1
        then {F X} | Rest
        else X | {Swap Rest (N - 1) F}
      end
    else
      nil
    end
  end

  fun {MatrixGet Matrix Row Col}
    {Nth {Nth Matrix (Row + 1)} (Col + 1)}
  end

  fun {MatrixSet Matrix Row Col X}
    {Swap Matrix (Row + 1) (fun {$ R} {Swap R (Col + 1) (fun {$ _} X end)} end)}
  end

  fun {MatrixNew Rows Cols}
    {Map {MakeList Rows} (fun {$ _} {MakeList Cols} end)}
  end

  fun {Multiply Chain}
    local
      N = {Length Chain}
      Aux = {FoldL {Range 0 (N - 1)} (fun {$ M I} {MatrixSet M I I {Nth Chain (I + 1)}} end) {MatrixNew N N}}
    in
      % Single matrix chain has zero cost
      for I in {Range 1 N} do % I: Length of subchain
        for J in {Range 0 (N - I)} do % I: Starting index of subchain
          local
            Best = 999999999 % Infinity
          in
            for K in {Range J (J + I)} do % K: Splitting point of subchain
              local
                % Multiply subchains at splitting point
                [LCost LName LRow LCol] = {MatrixGet Aux J K}
                [RCost RName _    RCol] = {MatrixGet Aux (K + 1) (J + I)}
                Cost = LCost + RCost + LRow * LCol * RCol
                Var = "(" # LName # RName # ")"
              in
                % Pick the best one
                if Cost < Best then
                  Best = Cost
                  Aux = {MatrixSet Aux J (J + I) [Cost Var LRow RCol]}
                end
              end
            end
          end
        end
      end

      Z = {MatrixGet Aux 0 (N - 1)}
      Z
    end
  end

  Z = {Multiply
    [[0 "A" 10 20]
     [0 "B" 20 30]
     [0 "C" 30 40]]}
  {System.showInfo Z}
end
