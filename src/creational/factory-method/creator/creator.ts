import { ITransport, SeaTransport, RoadTransport } from '../product/product';

abstract class TransportFactory {
    private transportList: ITransport[] = [];
    private id: number = 0;

    registerTransport(): ITransport {
        const transport = this.createTransport(this.id);

        this.transportList.push(transport);

        this.id++;

        return transport;
    }

    checkTransportCount(): number {
        console.log(this.transportList);

        return this.transportList.length;
    }

    abstract createTransport(id: number): ITransport
}

class SeaTransportFactory extends TransportFactory {
    createTransport(id: number): ITransport {
        return new SeaTransport(id);
    }
}

class RoadTransportFactory extends TransportFactory {
    createTransport(id: number): ITransport {
        return new RoadTransport(id);
    }
}

export { TransportFactory, SeaTransportFactory, RoadTransportFactory }
