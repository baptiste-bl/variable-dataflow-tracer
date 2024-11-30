# Fonction pour transformer le texte
def TransformText(text)
  text = text.upcase # Convertir en majuscules
  prefix = "Prefix: "
  AddPrefix(modified_text, prefix)
end

# Fonction pour ajouter un préfixe
def AddPrefix(text, prefix)
  prefix + text
end

# Fonction de test
def test
  input_text = "Hello, World!"
  result = TransformText(input_text)
  puts result
end

# Fonction principale
test
