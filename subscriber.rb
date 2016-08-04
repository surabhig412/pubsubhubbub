require 'sinatra'
require 'net/http'
require 'net/https'
require 'uri'
require 'json'
# set :protection, :except => :path_traversal


# get '/verify' do
#   puts "In request get"
#   puts params["challenge"]
#   content_type :json
#     { :challenge => params["challenge"], :topic => params["topic"]}.to_json
# end
#
# post '/publish' do
#   puts "Receiving updates from the publisher "
#   req = request.body.read
#   req = req.to_s
#     puts req
#   req =req.to_json
#   puts req
#   puts "***********IN HANDLER SINATRA*****************"
# end


get '/get' do
  puts "In request get method"
  content_type :json
    { :challenge => "VICTORY"}.to_json
end

post '/' do
  puts "In request post method"
	#request_payload = JSON.parse request.body.read
puts params
  
end
