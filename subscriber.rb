require 'sinatra'
require 'net/http'
require 'net/https'
require 'uri'


get '/verify' do
  puts "In subscription server verify method"
  redirect("http://localhost:8080/verify_intent?topic=#{params["topic"]}&challenge=#{params["challenge"]}")
end
