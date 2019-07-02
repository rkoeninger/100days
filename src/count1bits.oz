functor
import
  System
define
  proc {Count1Bits Value}
    if Value == 0 then
      1 + {Count1Bits (& Value (Value - 1))}
    else
      0
    end
  end

  {System.showInfo {Count1Bits }}
end
