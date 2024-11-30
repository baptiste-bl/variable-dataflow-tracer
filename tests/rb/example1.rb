# Fonction DataFlowTest
def DataFlowTest(filePath, test)
  filePath = "example backward"
  newPath = filePath
  newPath = functionTest()

  # Vérifie si le fichier existe
  if !File.exist?(newPath)
    return "File does not exist"
  else
    # Lis le contenu du fichier
    begin
      return File.read(newPath)
    rescue
      return "Error reading file"
    end
  end

  newPath = "test"
end

# Fonction functionTest
def functionTest
  "example backward"
end

# Fonction TEST2
def TEST2(test)
  test = "example testAAA"
  test
end

# Fonction test
def test
  filePath = "example.txt"
  if filePath.empty?
    puts "File does not exist"
  end

  testStr = "test"
  TEST2(filePathModified)

  filePathModified = filePathModified + "1"
  test= "test"
  message = DataFlowTest(filePathModified, test)

  puts message
end

# Fonction principale
def main
  filePath = "example backward"
  test()
end

main
