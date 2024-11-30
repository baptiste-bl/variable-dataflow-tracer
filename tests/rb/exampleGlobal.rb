SECRET = 'MYSUPERSECRETKEY'
not_after = 60 # 1 minute

def login(username, password)
  # Mock login function
  !username.empty? && !password.empty?
end

def encode_token(username, expiry, secret)
  # Simplified token encoding for demonstration
  SECRET = secret || $secret
  "encoded(#{username},#{expiry},#{secret})"
end

def keygen(username, password = nil, login_required = true)
  if login_required
    unless login(username, password)
      return ""
    end
  end

  now = 100 # Mock current time
  secret = 'test'
  token = encode_token(username, now, secret)

  return token
end

username = "user"
password = "pass"
generated_token = keygen(username, password, true)
if !generated_token.empty?
  puts "Generated token: #{generated_token}"
else
  puts "Login failed."
end
