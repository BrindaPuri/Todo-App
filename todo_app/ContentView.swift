//
//  ContentView.swift
//  todo_app
//
//  Created by Brinda Puri on 3/29/23.
//

import SwiftUI
import Combine

struct ContentView: View {
    @ObservedObject var taskStore = TaskStore()
    @State var supernewTask: String = ""
    var searchBar: some View{
        HStack{
            TextField("Insert New Task", text:self.$supernewTask)
            Button(action: self.addNewText, label:{Text("Add new")})
        }
    }
    func addNewText(){
        taskStore.tasks.append(Task(id: String(taskStore.tasks.count+1),NewTask: supernewTask))
        self.supernewTask=""
    }
    func move(from source : IndexSet, to destination : Int){
        taskStore.tasks.move(fromOffsets: source, toOffset: destination)
        
    }
    func del(at offset:IndexSet){
        taskStore.tasks.remove(atOffsets:offset)
    }
    
    var body: some View {
        NavigationView{
            VStack {
                searchBar.padding()
                List{
                    ForEach(self.taskStore.tasks){
                        task in
                        Text(task.NewTask)
                    }.onMove(perform: self.move)
                        .onDelete(perform: self.del)
                }.navigationTitle("To do List")
            }.navigationBarItems(trailing: EditButton())
        }
    }
}



struct ContentView_Previews: PreviewProvider {
    static var previews: some View {
        ContentView()
    }
}
