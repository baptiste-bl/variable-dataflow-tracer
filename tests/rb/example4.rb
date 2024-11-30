# Fonction principale
def main
  # numbers n'est pas défini, donc simuler l'initialisation à 1
  numbers = 1
  CrawlFromLine(numbers)
end

# CrawlFromLine analyse un fichier (représenté ici simplement par des lignes numérotées)
def CrawlFromLine(line)
  puts "Analyzing line: #{line}"

  # Condition pour simuler une fin de fichier à la ligne 10
  if line > 10
    puts "End of file reached."
    return
  end

  # Simuler un appel récursif à une fonction interne
  AnalyzeFunction(line + 1)
end

# AnalyzeFunction simule l'analyse d'une fonction à partir de la ligne actuelle
def AnalyzeFunction(line)
  puts "Entering function at line: #{line}"

  # Simuler un appel récursif à CrawlFromLine
  CrawlFromLine(line + 1)
end

# Appel de la fonction principale
main
