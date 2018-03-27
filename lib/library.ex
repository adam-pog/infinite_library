defmodule Library do
  @mapping Application.fetch_env!(:library, :mapping)
  @inverse_mapping Application.fetch_env!(:library, :inverse_mapping)
  @offset Application.fetch_env!(:library, :offset)
  @shift 0
  @base 29
  
  def encode(list) do
    Enum.map(list, &(@mapping[&1]))
    |> vol
  end
  
  defp vol(list) do
    list
    |> Enum.reverse
    |> shift(@shift)
    |> Enum.with_index
    |> Enum.reduce(0, fn({num, index}, final) -> 
      new = rem(num + elem(@offset, index), @base)
      IO.puts "#{num} -> #{new}, i: #{index}"
      final + trunc(Math.pow(@base, index) * new)
    end)
  end
  
  def decode(num) do
    given_location = num |> Math.to_base(@base)
    List.duplicate(0, 10 - length(given_location)) ++ given_location
    |> Enum.reverse
    |> Enum.with_index
    |> Enum.map(fn({num, index}) -> 
      IO.puts "#{num}, i: #{index}"
      rem(num + (@base - elem(@offset, index)), @base)
    end)
    |> Enum.reverse
    |> Enum.map(&(@inverse_mapping[&1]))
  end

  defp shift(list, 0) do
     list
  end
  
  defp shift([h|t], amt) do
    shift(t ++ [h], amt - 1)
  end
end
