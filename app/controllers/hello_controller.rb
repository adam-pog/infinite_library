class HelloController < ApplicationController
  def hello; end

  def world
    # @text = [params[:thing][0] * 80] * 40
    text = [params[:thing][0]] * 3200
    @text = text.each_slice(80).map{ |slice| slice.join('') }.join("\n")

  end
end
