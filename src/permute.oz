functor
import
  System
define
  fun {FindPivot Values}
    local
      Xs = {Filter {List.number 1 ({Length Values} - 1) 1} (fun {$ X} {Value.'<' {Nth Values X} {Nth Values (X + 1)}} end)}
    in
      if nil == Xs
        then 1
        else {Nth Xs 1}
      end
    end
  end

  fun {Permute Values}
    local
      N = {Length Values}
    in
      0
    end
  end

  {System.print {FindPivot [6 7 8 9]}}
end
