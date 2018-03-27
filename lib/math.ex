defmodule Math do
  require Integer

  def pow(_, 0), do: 1
  def pow(x, n) when Integer.is_odd(n), do: x * pow(x, n - 1)
  def pow(x, n) do
    result = pow(x, div(n, 2))
    result * result
  end
  
  def to_base(num, _) when num <= 29 do
    [num]
  end  
  
  def to_base(num, base) do
    to_base(div(num, base), base) ++ [rem(num, base)]
  end
end