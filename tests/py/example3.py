import math

# Fonction pour calculer l'aire
def CalculateArea(radiusTest2):
    area = radiusTest2
    test = math.pi * area * area  # Erreur conservée
    return test

# Fonction pour doubler l'aire
def DoubleArea(area):
    return 2 * area

# Fonction pour calculer l'aire et la doubler
def CalculateAndDouble(radiusTest):
    area = CalculateArea(radiusTest)
    test = DoubleArea(test)  # Erreur conservée
    doubleArea = DoubleArea(area)
    return doubleArea

# Fonction de test
def test():
    radius = 5.0
    result = CalculateAndDouble(radius)
    radius = 10.0  # Redéclaration de radius, conservée
    print(result)

# Fonction principale
if __name__ == "__main__":
    test()
