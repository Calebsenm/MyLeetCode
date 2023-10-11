

"""
1,2,3,4,5

3  3  5  6
4  5  6  7
5  6  7  8 
6  7  9  9 

"""  

#funcion que retorna los valores de los indices

def Fnumero(nums, target):
    numeroIndice = []
 
    for i in range(len(nums)):
        for j in range(len(nums)):
            if nums[j]  == nums[i]:
                if  len(nums) == 1:
                    continue
            if nums[i] + nums[j] == target:
                numeroIndice.append([i,j])

    for i in range(len(numeroIndice)):
        numeroIndice[i].sort()

    num1 = len(numeroIndice) -1
    
    print(numeroIndice)

    iterator2 = 0
    while iterator2   <  num1:
        iterator  =  0
        while iterator  < num1 :
            if numeroIndice[iterator2] ==  numeroIndice[iterator]:
                numeroIndice.pop(iterator2)
            iterator += 1    
            num1 = len(numeroIndice)  -1    
        iterator2  +=  1

    num2 =   []

    for i in numeroIndice:
        num2 = i
        break
    return num2

ListaNumeros = []
cantidadNumeros = int(input("Numeros en lista "))

for i in range(cantidadNumeros):
   numero = int(input(" numero => "))
   ListaNumeros.append(numero)


numeroBuscar = int(input("ingresa Un numero "))
print( ListaNumeros)

print(Fnumero(ListaNumeros , numeroBuscar ))















