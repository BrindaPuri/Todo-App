//
//  data_storage.swift
//  todo_app
//
//  Created by Brinda Puri on 3/29/23.
//

import Foundation
import SwiftUI
import Combine

                //stable identity
struct Task : Identifiable{
    var id=String()
    var NewTask=String()
}

class TaskStore : ObservableObject{
    @Published var tasks = [Task]()   
}
