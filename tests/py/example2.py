# Fonction pour transformer le texte
def TransformText(text):
    text = text.upper()  # Convertir en majuscules
    prefix = "Prefix: "
    return AddPrefix(modifiedText, prefix)

# Fonction pour ajouter un préfixe
def AddPrefix(text, prefix):
    return prefix + text

# Fonction de test
def test():
    inputText = "Hello, World!"
    result = TransformText(inputText)
    print(result)

# Fonction principale
if __name__ == "__main__":
    test()
