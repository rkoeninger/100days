functor
import
  System
define
  proc {Hanoi_H Height From To Through}
    if Height > 0 then
      {Hanoi_H (Height - 1) From Through To}
      {System.showInfo From#" => "#To}
      {Hanoi_H (Height - 1) Through To From}
    end
  end

  proc {Hanoi Height}
    {Hanoi_H Height "left" "right" "middle"}
  end

  {Hanoi 1}
  {System.showInfo ""}
  {Hanoi 2}
  {System.showInfo ""}
  {Hanoi 3}
  {System.showInfo ""}
  {Hanoi 4}
  {System.showInfo ""}
end
