import { Injectable } from '@angular/core';
import { GridServiceClient } from '../pb/grid_pb_service';
import { GridNode } from '../pb/grid_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';

@Injectable({
  providedIn: 'root',
})
export class HelloServiceService {
  private client: GridServiceClient;

  constructor() {
    this.client = new GridServiceClient('http://localhost:9090');
  }

  public async getGreeting(): Promise<GridNode[]> {
    const greetingPromise: Promise<GridNode[]> = new Promise((resolve) => {
      this.client.getGrid(new google_protobuf_empty_pb.Empty(), (err, res) =>
        resolve(res?.getNodesList() ?? [])
      );
    });

    return greetingPromise;
  }

  // public async getGreeting(): Promise<string> {
  //   const req = new GetNodeRequest();
  //   let greetingPromise: Promise<string> = new Promise((resolve) => {
  //     req.setName('Dmytro');
  //     this.client.sayHello(req, (err, res) =>
  //       resolve(res?.getGreeting() ?? '')
  //     );
  //   });

  //   return greetingPromise;
  // }
}
